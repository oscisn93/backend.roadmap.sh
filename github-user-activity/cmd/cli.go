package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/cmd/client"
)

type CLI struct {
	Name        string
	Description string
	Args        []string
}

func (cli *CLI) parseArgs() error {
	fmt.Print(os.Args)

	if len(os.Args) < 1 {
		os.Exit(1)
	}

	cli.Args = os.Args

	return nil
}

func (cli *CLI) createUserRequest() (client.PublicUserEvents, error) {
	return nil, nil
}

func (cli *CLI) displayEvents(_ client.PublicUserEvents) {
	fmt.Println("")
}

func (cli *CLI) Execute() error {
	err := cli.parseArgs()

	if err != nil {
		return err
	}

	events, err := cli.createUserRequest()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cli.displayEvents(events)

	return nil
}

// "github-activity",
// "A program to fetch user activiy from GitHub"
// `GitHub Activity is a CLI tool that allows us to query
//         user activity events from the GitHub REST API,
//         right from the command terminal.`
