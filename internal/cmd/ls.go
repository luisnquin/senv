package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/gookit/color"
	"github.com/luisnquin/senv/internal/core"
)

var rxSenvDotEnvComment = regexp.MustCompile("[\\#\\_]+")

func Ls(currentDir string) error {
	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	switcher := NewSwitcher(settings)

	active, err := switcher.GetActiveEnvironment()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	envNames := make([]string, len(settings.Environments))

	for i, e := range settings.Environments {
		envNames[i] = e.Name
	}

	sort.Strings(envNames)

	namePrinter := color.New(color.HiGreen, color.OpItalic) // color.HEX("#a4d43f")
	activePrinter := color.New(color.LightCyan, color.Bold) // color.HEX("#84dde0")

	fmt.Fprintf(os.Stdout, "%s:\n", settings.SourceFilePath)

	for _, name := range envNames {
		activeLabel := ""

		if name == active {
			activeLabel = activePrinter.Sprint("(active)")
		}

		fmt.Fprintf(os.Stdout, "- %s %s\n", namePrinter.Sprint(name), activeLabel)
	}

	return nil
}
