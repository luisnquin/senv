package cmd

import (
	"os"

	"github.com/luisnquin/senv/internal/core"
)

func Check() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if !core.HasConfigFiles(currentDir) {
		os.Exit(1)
	}

	return nil
}
