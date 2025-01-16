package githubapi

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/internal/cache/libsql"
)

type UserRequest struct {
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
	Etag      string `json:"etag"`
	RateLimit int    `json:"rateLimit"`
}

type ApiConfig struct {
	Username string
	Token    string
	Url      string
}

func getPublicToken(envFile string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Could not load environmnet variables from file:", envFile)
	}
	token := os.Getenv("GITHUB_PUBLIC_API_TOKEN")
	return token
}

func createAppConfig(username string) ApiConfig {
	file := ".env"
	tokenString := getPublicToken(file)
	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
	config := ApiConfig{
		Username: username,
		Token:    tokenString,
		Url:      url,
	}
	return config
}

func CreateUserRequest(username string, q *libsql.Queries) {
	userRequests, err := q.GetUserRequests()
}
