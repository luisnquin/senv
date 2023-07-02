package cmd

import (
	"strings"

	"github.com/luisnquin/senv/env"
	"github.com/luisnquin/senv/log"
)

func SwitchTo(currentDir, envToSwitch string) error {
	envToSwitch = strings.TrimSpace(envToSwitch)

	workDir, err := env.ResolveUsableWorkDir(currentDir)
	if err != nil {
		return err
	}

	environments, err := env.LoadEnvironments(workDir)
	if err != nil {
		return err
	}

	if err := switchDotEnvFileFromName(workDir, environments, envToSwitch); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", envToSwitch)

	return nil
}
