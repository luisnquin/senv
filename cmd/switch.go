package cmd

import (
	"errors"
	"os"

	"github.com/luisnquin/senv/internal/env"
	"github.com/luisnquin/senv/internal/log"
	"github.com/luisnquin/senv/internal/prompt"
	"github.com/samber/lo"
)

// Creates a prompt selector that allows the user to select the environment to switch.
func Switch(currentDir string) error {
	if !env.HasUsableWorkDir(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`") // or `.env` files")
	}

	preferences, err := env.LoadUserPreferences()
	if err != nil {
		return err
	}

	envNames := make([]string, len(preferences.Environments))

	for i, env := range preferences.Environments {
		envNames[i] = env.Name
	}

	selected, ok := prompt.ListSelector("Select an environment", envNames)
	if !ok {
		os.Exit(1)
	}

	if err := switchDotEnvFileFromName(preferences, selected); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", selected)

	return nil
}

func switchDotEnvFileFromName(preferences *env.UserPreferences, envToSwitch string) error {
	environment, ok := lo.Find(preferences.Environments, func(e env.Environment) bool {
		return e.Name == envToSwitch
	})
	if !ok {
		return errors.New("environment not found")
	}

	dotEnvData, err := env.GenerateDotEnv(environment)
	if err != nil {
		return err
	}

	envFilePath, err := preferences.GetEnvFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(envFilePath, dotEnvData, os.ModePerm); err != nil {
		return err
	}

	return nil
}
