package main

import "github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/internal/api"

type GitHubClient struct {
  restAPI *api.API
}

func (ghc *GitHubClient)New () {
  ghc = &GitHubClient{}
}
