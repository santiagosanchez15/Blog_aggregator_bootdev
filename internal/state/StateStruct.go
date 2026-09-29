package state

import (
	_ "github.com/lib/pq"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/config"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/database"
)

type State struct {
	Pconfig *config.Config
	Db      *database.Queries
}
