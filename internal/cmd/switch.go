package cmd

import (
	"errors"
	"os"

	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/log"
	"github.com/luisnquin/senv/internal/prompt"
	"github.com/samber/lo"
)

// Creates a prompt selector that allows the user to select the environment to switch.
func Switch(currentDir string) error {
	if !core.WorkDirHasProgramFiles(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`")
	}

	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	var activeEnv string

	if envFilePath, err := settings.GetEnvFilePath(); err == nil {
		activeEnv = getActiveEnvironment(envFilePath)
	}

	envNames := make([]string, len(settings.Environments))

	for i, env := range settings.Environments {
		envNames[i] = env.Name
	}

	selected, ok := prompt.ListSelector("Select an environment", envNames, activeEnv)
	if !ok {
		os.Exit(1)
	}

	if err := switchDotEnvFileFromName(settings, selected, settings.UseExportPrefix); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", selected)

	return nil
}

func switchDotEnvFileFromName(preferences *core.UserPreferences, envToSwitch string, useExportPrefix bool) error {
	environment, ok := lo.Find(preferences.Environments, func(e core.Environment) bool {
		return e.Name == envToSwitch
	})
	if !ok {
		return errors.New("environment not found")
	}

	dotEnvData, err := core.GenerateDotEnv(environment, useExportPrefix)
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
