package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/gookit/color"
	"github.com/luisnquin/senv/internal/core"
	"github.com/samber/lo"
)

var rxSenvDotEnvComment = regexp.MustCompile("[\\#\\_]+")

func Ls(currentDir string, raw bool) error {
	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	envFilePath, err := settings.GetEnvFilePath()
	if err != nil {
		return err
	}

	activeEnv := getActiveEnvironment(envFilePath)

	envNames := make([]string, len(settings.Environments))

	for i, e := range settings.Environments {
		envNames[i] = e.Name
	}

	sort.Strings(envNames)

	activeLabel := "(active)"
	if !raw {
		activePrinter := color.New(color.LightCyan, color.Bold) // color.HEX("#84dde0")
		activeLabel = activePrinter.Sprint(activeLabel)

		fmt.Fprintf(os.Stdout, "%s:\n", settings.SourceFilePath)
	}

	namePrinter := color.New(color.HiGreen, color.OpItalic) // color.HEX("#a4d43f")

	for _, name := range envNames {
		isCurrentEnv := name == activeEnv

		if !raw {
			name = namePrinter.Sprint(name)
		}

		fmt.Fprintf(os.Stdout, "- %s %s\n", name, lo.If(isCurrentEnv, activeLabel).Else(""))
	}

	return nil
}
