package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeConfigJson(c Config) error {
	/*Write config to json after done*/

	data, err := json.Marshal(c) //marshal data from config
	if err != nil {
		return fmt.Errorf("Error when coding data| Error: %w", err)
	}
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error when gettingn the filepath| Error %w", err)
	}
	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return fmt.Errorf("Error when writing the file %w", err)
	}

	return nil
}
