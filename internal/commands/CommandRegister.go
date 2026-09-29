package commands

import (
	"context"

	"time"

	"fmt"

	"github.com/google/uuid"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/database"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func Register(s *state.State, cmd Command) error {
	// takes state and cmd and register the user

	if len(cmd.Args) <= 0 {
		return fmt.Errorf("Args empty please provide a name")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	new_user, err := s.Db.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Error when creating user %w", err)
	}

	err = s.Pconfig.SetUser(new_user.Name)
	if err != nil {
		return fmt.Errorf("error when setting user %w", err)
	}
	fmt.Printf("User was created")

	return nil
}
