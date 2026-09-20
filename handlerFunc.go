package main

import (
	"fmt"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/commands"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func handlerLogin(s *state.State, cmd commands.Command) error {
	/*takes the command struct and uses the slice of string to get username and set up the cofug pointer*/

	if len(cmd.Args) <= 0 { // check if args is empty
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	name := cmd.Args[0] // get username

	s.Pconfig.SetUser(name)               // set username
	fmt.Printf("The user has been set\n") // print success

	return nil
}
