package config

import (
	"os"
	"projectname/internal/project/domain/configuration"
)

func determineBaseDir() (string, error) {
	if err := os.MkdirAll(configuration.DefaultBaseDir, 0755); err != nil {
		return "", err
	}

	return configuration.DefaultBaseDir, nil
}
