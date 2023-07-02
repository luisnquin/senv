package cmd

import (
	"errors"
	"os"
)

func Init(configFile []byte) error {
	const configFilePath = "./senv.yaml"

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
