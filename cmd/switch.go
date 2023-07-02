package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/luisnquin/senv/env"
	"github.com/luisnquin/senv/log"
	"github.com/luisnquin/senv/prompt"
	"github.com/samber/lo"
)

// Creates a prompt selector that allows the user to select the environment to switch.
func Switch(currentDir string) error {
	if !env.HasUsableWorkDir(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml` or `.env` files")
	}

	workDir, err := env.ResolveUsableWorkDir(currentDir)
	if err != nil {
		return err
	}

	environments, err := env.LoadEnvironments(workDir)
	if err != nil {
		return err
	}

	envNames := make([]string, len(environments))

	for i, env := range environments {
		envNames[i] = env.Name
	}

	selected, ok := prompt.ListSelector("Select an environment", envNames)
	if !ok {
		os.Exit(1)
	}

	if err := switchDotEnvFileFromName(workDir, environments, selected); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", selected)

	return nil
}

func switchDotEnvFileFromName(workDir string, envs []env.Environment, envToSwitch string) error {
	environment, ok := lo.Find(envs, func(e env.Environment) bool {
		return e.Name == envToSwitch
	})
	if !ok {
		return errors.New("environment not found")
	}

	dotEnv, err := env.GenerateDotEnv(environment)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(workDir, ".env"), dotEnv, os.ModePerm); err != nil {
		return err
	}

	return nil
}
