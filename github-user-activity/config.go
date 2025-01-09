package main

import "os"

func getPublicEventsToken(filename: String) {
  token := os.Getenv("GITHUB_PUBLIC_API_TOKEN")
  return token
}

type CliDependency struct {
  name: String
  api: {
    type: String;
    token: String;
  }
}

type CliConig struct {
  dependencies: CliDependency[]
}

var cliConfig :CliConig = {
  dependencies: {
    github: {
      api: {
        type: "PublicUserEvents",
        token: getPublicEventsToken() 
      }
    }
  }
  
}