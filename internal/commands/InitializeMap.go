package commands

import (
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func InitializeCommandsMap() Commands {
	/*Initializes the command struct and its map*/

	return Commands{MapCommands: make(map[string]func(*state.State, Command) error)}
}
