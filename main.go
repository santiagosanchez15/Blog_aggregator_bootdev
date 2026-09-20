package main

import (
	"fmt"

	"github.com/santiagosanchez15/Blog_aggregator_bootdev/internal/config"
)

func main() {

	first_config, err := config.Read() // get config and forget about the error
	if err != nil {                    // check if reading return wrong error
		return
	}
	name := "Santiago"
	first_config.SetUser(name) // set user and write to file
	new_config, err := config.Read()

	if err != nil { // check uif reading return error
		return
	}
	fmt.Printf("%v\n%v\n", new_config.DBURL, new_config.CurrentUserName)
}
