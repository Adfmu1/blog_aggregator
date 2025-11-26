package commands

import (
	"github.com/Adfmu1/blog_aggregator/internal/config"
)

type command struct {
	name      string
	arguments []string
}

type state struct {
	file *config.Config
}
