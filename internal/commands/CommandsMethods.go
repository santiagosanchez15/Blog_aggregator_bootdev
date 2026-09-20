package commands

import (
	"fmt"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func (c Commands) Run(s *state.State, cmd Command) error {
	/*Runs the command with the given state if exists*/

	if s == nil {
		return fmt.Errorf("Given state points to nothing please add a correct address to it")
	}

	err := c.MapCommands[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c Commands) Register(name string, f func(s *state.State, cmd Command) error) error {

	c.MapCommands[name] = f // add function to map and add the value to call the function
	return nil
}
