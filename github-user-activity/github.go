package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func GetPublicToken(filename string) string {
	err := godotenv.Load(filename)

	if err != nil {
		log.Fatal("Could not load environmnet variables from file:", filename)
	}

	token := os.Getenv("GITHUB_PUBLIC_API_TOKEN")

	return token
}

type Datastore interface {
	Get
	Set
}

type Client struct {
	*GitHubCache
}

type UserEntry struct {
	Filename string `json:"filename`
	Etag     string `json:"etag"`
}

type UserRequest struct {
	Username  string `json:"username`
	Time      int64  `json:"time"`
	RateLimit int    `json:"rateLimit"`
}

type GitHubCache struct {
	LastRequest UserRequest          `json:"lastRequest"`
	UserEntries map[string]UserEntry `json:"userEntries"`
}

func GitHubCacheClient(jsonFilename string) *Client {
	file, err := os.Open(jsonFilename)

	if err != nil {
		log.Fatal("Something went wrong with parsing the file at", jsonFilename)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)

	var ghCache GitHubCache
	err = json.Unmarshal(content, &ghCache)

	if err != nil {
		log.Fatal("Failed to create git hub cache.\n", err)
	}

	return &Client{&ghCache}
}

type GitHubCacheClientError struct {
	EmptyUsers        error
	UserNotFound      error
	NoPreviousRequest error
}

var ClientCacheErrors = GitHubCacheClientError{
	EmptyUsers:        errors.New("there are currently no users, therefore no rate limits can be applied"),
	UserNotFound:      errors.New("the user with that username was not found in the data store"),
	NoPreviousRequest: errors.New("there are no previous requests to this cache"),
}

func (c *Client) getUserEntries() (map[string]UserEntry, error) {
	size := len(c.UserEntries)
	if size == 0 {
		return nil, ClientCacheErrors.EmptyUsers
	}
	return c.UserEntries, nil
}

func (c *Client) GetUserEntry(username string) (*UserEntry, error) {
	userEntries, err := c.getUserEntries()

	if err != nil {
		log.Fatal("User cannot exist in the data store as it is empty")
	}

	user, ok := userEntries[username]

	if !ok {
		return nil, ClientCacheErrors.UserNotFound
	}

	return &user, nil
}

func (c *Client) GetLastRequest() (*UserRequest, error) {
	lastRequest := c.LastRequest
	if lastRequest.Time == 0 {
		return nil, ClientCacheErrors.NoPreviousRequest
	}
	return &lastRequest, nil
}

// TODO: complete this logic
func (c *Client) GetStoredUserEvents(username string) (PublicUserEvents, error) {
	return nil, nil
}

type Get interface {
	GetUserEntry(username string) (*UserEntry, error)
	GetLastRequest() (*UserRequest, error)
	GetStoredUserEvents(username string) (PublicUserEvents, error)
}

func (c *Client) AddUserRequest(username string, rateLimit int) {
	currentTime := time.Now().Unix()
	latestRequest := UserRequest{
		Username:  username,
		Time:      currentTime,
		RateLimit: rateLimit,
	}
	c.LastRequest = latestRequest
}

func (c *Client) writeFile(filename string, events PublicUserEvents) {
	jsonFile, error := os.Open(filename)

	if error != nil {
		log.Fatal("The file at", filename, "could not be found")
	}
	defer jsonFile.Close()

	data, error := json.MarshalIndent(events, "", "    ")

	if error != nil {
		log.Fatal("Could not parse the data from the 'PublicUserEvents' object")
	}
	_, writeError := jsonFile.Write(data)
	if writeError != nil {
		log.Fatal("Could not write the file for the user from the previous query")
	}
}

func (c *Client) SetUserEntry(username string, etag string, events PublicUserEvents) bool {
	userEntry, err := c.GetUserEntry(username)

	if err != nil {
		filename := fmt.Sprintf("./entries/%s.json", username)
		c.UserEntries[username] = UserEntry{
			Filename: filename,
			Etag:     etag,
		}
		c.writeFile(filename, events)
		return true
	}

	if userEntry.Etag == etag {
		return false
	}

	userEntry.Etag = etag

	c.writeFile(c.UserEntries[username].Filename, events)

	return true
}

type Set interface {
	AddUserRequest(username string, rateLimit int)
	SetUserEntry(username string, etag string) bool
}
