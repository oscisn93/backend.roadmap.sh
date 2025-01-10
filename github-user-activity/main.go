package main

import (
  "net/http"
  "fmt"
  "log"
  "io"

)

func fetchUserActivity(username string) {
  url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
  request, error := http.NewRequest("GET", url, nil)

  if error != nil {
    log.Fatal("Bad Request")
  }
  
  envFile := ".env"
  token := GetPublicEventsToken(envFile)
  tokenString := fmt.Sprintf("Bearer %s", token)

  request.Header.Set("Content-Type", "application/json")
  request.Header.Set("Authorization", tokenString)

  client := &http.Client{}

  response, error := client.Do(request)
  if error != nil {
    log.Fatalf("Something went wrong. %v", error)
  }
  defer response.Body.Close()

  body, error := io.ReadAll(response.Body)
  if error != nil {
    log.Fatalf("Something went wrong. %v", error)
  }

  fmt.Println("Response status:", response.Status)
  fmt.Println("Response body:", string(body))
}
