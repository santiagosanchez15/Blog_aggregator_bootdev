package commands

import (
	"context"
	"fmt"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func Reset(s *state.State, cmd Command) error {

	err := s.Db.ResetTable(context.Background()) // reset table

	if err != nil {
		return err
	}

	fmt.Printf("User deleted\n")

	return nil
}
