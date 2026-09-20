package config

import "fmt"

func (c *Config) SetUser(name string) error {
	/*Writes config struct to json after filling up the username*/
	c.CurrentUserName = name   //set name
	err := writeConfigJson(*c) // write to file
	if err != nil {
		return fmt.Errorf("Error when writing the file %w", err)
	}

	return nil
}
