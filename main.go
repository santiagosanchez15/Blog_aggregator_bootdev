package main

import (
	"fmt"

	"os"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/commands"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/config"
	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/state"
)

func main() {
	// Initialize the state
	s := state.State{} // Create state struct

	c, err := config.Read() // get config and forget about the error
	if err != nil {         // check if reading return wrong error
		fmt.Printf("Error given : %v\n", err)
		os.Exit(2)
	}
	s.Pconfig = &c // get pointer to config struct

	// Initialize the commands
	m := commands.InitializeCommandsMap()   //create map of commands
	err = m.Register("login", handlerLogin) // Add the command to the commands map
	if err != nil {                         // check if register got an error
		fmt.Printf("Error given : %v\n", err)
		os.Exit(2)
	}

	// Create command and get args
	allArgs := os.Args               // get cli arguments
	args, err := getCliArgs(allArgs) // split and get useful args
	if err != nil {                  // check if error
		fmt.Printf("Error given : %v\n", err)
		os.Exit(1)
	}

	cmd := commands.CreateCommand(args[0], args[1:]) // create the comamnd

	err = m.Run(&s, cmd)
	if err != nil {
		fmt.Printf("Error given : %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
