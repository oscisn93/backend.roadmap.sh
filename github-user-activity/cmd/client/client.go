package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var envFile = ".env"

func FetchPublicUserEvents(username string) PublicUserEvents {
	token := GetPublicToken(envFile)
	url := fmt.Sprintf("https://github.com/api/users/%s/events", username)
	request, error := http.NewRequest("GET", url, nil)

	if error != nil {
		log.Fatal("Bad Request")
	}

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
