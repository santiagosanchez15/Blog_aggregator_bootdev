package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	// Function that reads jsoon and return Config struct with error

	path, err := getConfigFilePath() // get the file path
	if err != nil {
		return Config{}, fmt.Errorf("Error when getting the file path| Error:%w", err) // return error
	}
	body, err := os.ReadFile(path) // get the body after reading the file
	if err != nil {
		return Config{}, fmt.Errorf("Error when reading the file| error:%w", err) // raise error if error
	}

	new_config := Config{} // create holder

	err = json.Unmarshal(body, &new_config) // decode and get error if error
	if err != nil {
		return Config{}, fmt.Errorf("Error when decoding the file and passing to struct| %w", err) //return error
	}

	return new_config, nil //return new config

}
