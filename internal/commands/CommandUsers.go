package commands

import (
	"context"
	"fmt"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func Users(s *state.State, cmd Command) error {
	// Returns list with current users, list current for which session is initialized

	names, err := s.Db.GetUsers(context.Background())

	if err != nil {
		return err
	}

	for _, name := range names {
		fmt.Printf("* %s", name)
		if s.Pconfig.CurrentUserName == name {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")

	}

	return nil
}
