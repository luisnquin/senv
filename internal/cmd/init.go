package cmd

import (
	"errors"
	"os"

	"github.com/luisnquin/senv/internal/assets"
)

func Init() error {
	configFile := assets.GetExampleConfig()
	configFilePath := "./senv.yaml"

	_, err := os.Stat(configFilePath)
	if err == nil {
		return errors.New("configuration file already exists in current directory")
	}

	err = os.WriteFile(configFilePath, configFile, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
