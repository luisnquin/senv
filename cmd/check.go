package cmd

import (
	"os"

	"github.com/luisnquin/senv/env"
)

func Check() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if !env.HasUsableWorkDir(currentDir) {
		os.Exit(1)
	}

	return nil
}
