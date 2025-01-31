package cache

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/oscisn93/backend.roadmap.sh/tree/main/github-user-activity/cmd/cache/libsql"
	_ "github.com/tursodatabase/go-libsql"
)

var dbName string = "file:./local.db"

type Cache struct {
	Ctx    context.Context
	Client *libsql.Queries
}

func New() *Cache {
	cache := &Cache{}
	cache.Ctx = context.Background()
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	cache.Client = libsql.New(db)
	return cache
}
