package cmd

import (
	"bufio"
	"bytes"
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

	namePrinter := color.New(color.HiGreen, color.OpItalic) // color.HEX("#a4d43f")
	activePrinter := color.New(color.LightCyan, color.Bold) // color.HEX("#84dde0")

	fmt.Fprintf(os.Stdout, "%s:\n", settings.SourceFilePath)

	for _, name := range envNames {
		activeLabel := ""

		if name == activeEnv {
			activeLabel = activePrinter.Sprint("(active)")
		}

		fmt.Fprintf(os.Stdout, "- %s %s\n", namePrinter.Sprint(name), activeLabel)
	}

	return nil
}

// Returns an empty string if not found.
func getActiveEnvironment(envFilePath string) string {
	data, err := os.ReadFile(envFilePath)
	if err != nil {
		return ""
	}

	s := bufio.NewScanner(bytes.NewReader(data))

	for s.Scan() {
		line := s.Text()

		if len(rxSenvDotEnvComment.FindAllString(line, -1)) == 2 {
			return rxSenvDotEnvComment.ReplaceAllString(line, "")
		}
	}

	return ""
}
