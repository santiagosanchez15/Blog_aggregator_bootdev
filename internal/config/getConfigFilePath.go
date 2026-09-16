package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json" // not exportable

func getConfigFilePath() (string, error) {
	/*Gets the config file path to avoid joining the string*/

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error raised when getting userHomeDir| Error: %w", err)
	}
	path := filepath.Join(homeDir, configFileName)
	return path, nil
}
