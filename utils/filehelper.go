package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetAbsolutePath(relativePath string) (string, error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(absolutePath)
	if os.IsNotExist(err) {
		return "", errors.New(err.Error()[5:])
	}
	if !info.IsDir() {
		return "", errors.New(absolutePath + " is not a directory")
	}
	return absolutePath, nil
}
