package main

import (
	"fmt"
)

func getCliArgs(userArgs []string) ([]string, error) {
	/*Gets user args when running the program and returns a slice of the args split after the name*/

	if len(userArgs) < 2 {
		return nil, fmt.Errorf("The len of the arguments is less than two please add more") // return empty array plus error message
	}

	return userArgs[1:], nil
}
