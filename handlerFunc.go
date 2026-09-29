package main

import (
	"fmt"

	"context"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/commands"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func handlerLogin(s *state.State, cmd commands.Command) error {
	/*takes the command struct and uses the slice of string to get username and set up the cofug pointer*/

	if len(cmd.Args) <= 0 { // check if args is empty
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	name := cmd.Args[0] // get username
	username, err := s.Db.GetUser(context.Background(), name)

	if err != nil {
		return fmt.Errorf("User doesnt exist| error %w", err)
	}

	err = s.Pconfig.SetUser(username) // set username
	if err != nil {
		return fmt.Errorf("Error when setting user %w", err)
	}
	fmt.Printf("The user has been set\n") // print success

	return nil
}
