package main

import (
  "net/http"
  "fmt"
)

func fetchUserActivity(username: String) {
  url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
  response, error := http.Get()
}
