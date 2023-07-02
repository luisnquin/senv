package cmd

import (
	"strings"

	"github.com/luisnquin/senv/env"
	"github.com/luisnquin/senv/log"
)

func SwitchTo(currentDir, envToSwitch string) error {
	if !env.HasUsableWorkDir(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`") // or `.env` files")
	}

	envToSwitch = strings.TrimSpace(envToSwitch)

	settings, err := env.LoadUserPreferences()
	if err != nil {
		return err
	}

	if err := switchDotEnvFileFromName(settings, envToSwitch); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", envToSwitch)

	return nil
}
