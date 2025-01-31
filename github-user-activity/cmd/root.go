package cmd

import (
	"os"
)

var cli = &CLI{
	Name: "github-activity",
	Description: `GitHub Activity is a CLI tool that allows us to query 
         user activity events from the GitHub REST API, right 
         from the command terminal.`,
	Args: []string{"github-activity"},
}

func Execute() {
	err := cli.Execute()

	if err != nil {
		os.Exit(1)
	}
}
