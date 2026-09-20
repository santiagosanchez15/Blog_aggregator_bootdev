package commands

import (
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

type Command struct { // type command single element
	Name string
	Args []string
}

type Commands struct { // type commands will contain multiple Command structs
	MapCommands map[string]func(*state.State, Command) error
}
