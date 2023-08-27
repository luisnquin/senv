package test

import (
	"fmt"
	"os"
	"path/filepath"
)

func Mkdir() (string, error) {
	if err := os.MkdirAll("/tmp/senv", os.ModePerm); err != nil {
		return "", fmt.Errorf("unable to create /tmp/senv directory: %v", err)
	}

	return os.MkdirTemp("/tmp/senv", "ls-test-*")
}

func WriteTo(dir, file string, data []byte) error {
	path := filepath.Join(dir, file)

	return os.WriteFile(path, data, os.ModePerm)
}
