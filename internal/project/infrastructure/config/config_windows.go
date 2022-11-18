//go:build windows
// +build windows

package config

import (
	"os"
	"path/filepath"
	"projectname/internal/project/domain/configuration"
	"strings"
)

func determineBaseDir() (string, error) {
	exe, err := os.Executable()

	if err != nil {
		return "", err
	}

	dir := strings.TrimSuffix(filepath.Dir(exe), string(os.PathSeparator)+configuration.DefaultDirBin)

	if err = os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}
