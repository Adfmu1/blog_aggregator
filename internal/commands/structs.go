package commands

import (
	"github.com/Adfmu1/blog_aggregator/internal/config"
	"github.com/Adfmu1/blog_aggregator/internal/database"
)

type Commands struct {
	RegisteredComms map[string]func(*State, Command) error
}

type Command struct {
	Name      string
	Arguments []string
}

type State struct {
	Database *database.Queries
	File *config.Config
}
