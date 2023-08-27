package cmd

import (
	"fmt"
	"os"

	"github.com/luisnquin/senv/internal/core"
	"github.com/samber/lo"
)

func Out() error {
	settings, err := core.LoadUserPreferences()
	if err != nil {
		os.Exit(1)
	}

	switcher := NewSwitcher(settings)

	activeEnv, err := switcher.GetActiveEnvironment()
	if err != nil {
		os.Exit(1)
	}

	activeBetweenEnvsFn := func(env core.Environment) bool { return env.Name == activeEnv }

	if !lo.SomeBy(settings.Environments, activeBetweenEnvsFn) {
		os.Exit(1)
	}

	_, err = fmt.Fprintln(os.Stdout, activeEnv)

	return err
}
