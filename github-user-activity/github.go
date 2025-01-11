package main

import (
	"log"
	"os"

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
