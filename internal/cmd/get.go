package cmd

import (
	"fmt"
	"os"

	"github.com/luisnquin/senv/internal/core"
	"github.com/samber/lo"
)

func GetEnv() error {
	settings, err := core.LoadUserPreferences()
	if err != nil {
		os.Exit(1)
	}

	envFilePath, err := settings.GetEnvFilePath()
	if err != nil {
		os.Exit(1)
	}

	activeEnv := getActiveEnvironment(envFilePath)
	if activeEnv == "" {
		os.Exit(1)
	}

	activeBetweenEnvsFn := func(env core.EnvironmentDefinition) bool { return env.Name == activeEnv }

	if !lo.SomeBy(settings.Environments, activeBetweenEnvsFn) {
		os.Exit(1)
	}

	_, err = fmt.Fprintln(os.Stdout, activeEnv)

	return err
}
