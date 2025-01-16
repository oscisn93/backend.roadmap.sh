package api

import (
  "log"

	"github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/githubapi"
  "github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/internal/cache"
	"github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/internal/cache/libsql"
)

type EventsAPI interface {
  getEvents() githubapi.PublicUserEvents 
  getEvent(id string) githubapi.Event
  addEvent(libsql.AddEventParams) githubapi.Event
  updateEvent(id string, params libsql.UpdateEventParams) githubapi.Event
  deleteEvent(id string) bool
}

type API struct {
  eventsAPI *EventsAPI
}

func GetPublicEvents(c *cache.Cache) {
  events, err := c.Client.GetEvents(c.Ctx)
  if err != nil {
    log.Fatal("Cannot get public events")
  }
  publicEvents := [30]githubapi.PublicUserEvents{}
  for _, event := range events {
    actor, err := c.Client.GetActor(c.Ctx, event.ActorID)
    if err != nil {

    }

  }
}

