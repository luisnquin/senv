package cmd

import (
	"strings"

	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/log"
)

func SwitchTo(currentDir, envToSwitch string) error {
	if !core.HasConfigFiles(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`") // or `.env` files")
	}

	envToSwitch = strings.TrimSpace(envToSwitch)

	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	if err := switchDotEnvFileFromName(settings, envToSwitch, settings.UseExportPrefix); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", envToSwitch)

	return nil
}
