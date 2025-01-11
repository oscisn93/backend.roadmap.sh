package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GitHubActivityCLIConfig struct {
	Username string
	Token    string
	Url      string
}

func CreateCliConfig(username string) GitHubActivityCLIConfig {
	file := ".env"
	tokenString := GetPublicToken(file)
	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
	config := GitHubActivityCLIConfig{
		Username: username,
		Token:    tokenString,
		Url:      url,
	}
	return config
}

func fetchUserActivity(config GitHubActivityCLIConfig) PublicUserEvents {
	request, error := http.NewRequest("GET", config.Url, nil)

	if error != nil {
		log.Fatal("Bad Request")
	}

	var envFile = ".env"
	token := GetPublicToken(envFile)
	tokenString := fmt.Sprintf("Bearer %s", token)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", tokenString)
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		log.Fatalf("Something went wrong. %v", error)
	}

	defer response.Body.Close()

	etag := response.Header.Get("etag")

	if etag == "" {
		log.Fatal("No header called 'ETAG' exists on the response")
	}

	fmt.Println("ETAG value:", etag)

	rateLimitRemaining := response.Header.Get("x-ratelimit-remaining")

	if rateLimitRemaining == "" {
		log.Fatal("No header 'x-ratelimit-remaining' exists on the response")
	}

	fmt.Println("Rate Limit Remaining: ", rateLimitRemaining)

	body, error := io.ReadAll(response.Body)

	if error != nil {
		log.Fatalf("Something went wrong. %v", error)
	}

	events, error := UnmarshalPublicUserEvents(body)

	if error != nil {
		log.Fatal("Unable to unmarhsal public user events. Something went wrong")
	}

	return events

}

func GetUserEvents(config GitHubActivityCLIConfig) PublicUserEvents {
	ghCache := GitHubCacheClient("github-cache.json")
	user, error := ghCache.GetUserEntry(config.Username)
	if error != nil {
		fetchUserActivity(config)
	}
	etag := user.Etag

}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must specify a github username after the command.")
	}
	usernameArg := os.Args[1]
	config := CreateCliConfig(usernameArg)
	var events = GetUserEvents(config)

	data, err := json.MarshalIndent(events, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PublicEvents for GitHub user", config.Username, "\n", string(data))
}
