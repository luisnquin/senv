package cmd

import (
	"os"

	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/log"
	"github.com/luisnquin/senv/internal/prompt"
)

// Creates a prompt selector that allows the user to select the environment to switch.
func Switch(currentDir string) error {
	if !core.HasConfigFiles(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`")
	}

	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	switcher := NewSwitcher(settings)

	activeEnv, err := switcher.GetActiveEnvironment()
	if err != nil {
		return err
	}

	envNames := make([]string, len(settings.Environments))

	for i, env := range settings.Environments {
		envNames[i] = env.Name
	}

	selected, ok := prompt.ListSelector("Select an environment", envNames, activeEnv)
	if !ok {
		os.Exit(1)
	}

	if err := switcher.Switch(selected); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", selected)

	return nil
}
