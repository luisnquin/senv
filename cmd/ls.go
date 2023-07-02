package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gookit/color"
	"github.com/luisnquin/senv/env"
)

var rxSenvDotEnvComment = regexp.MustCompile("[\\#\\_]+")

func Ls(currentDir string) error {
	workDir, err := env.ResolveUsableWorkDir(currentDir)
	if err != nil {
		return err
	}

	environments, err := env.LoadEnvironments(workDir)
	if err != nil {
		return err
	}

	activeEnv := getActiveEnvironment(workDir)

	for _, e := range environments {
		active := e.Name == activeEnv

		if active {
			color.HEX("#7de8e8").Printf("(on)  %s\n", e.Name)
		} else {
			fmt.Fprintf(os.Stdout, "(off) %s\n", e.Name)
		}

	}

	return nil
}

// Returns an empty string if not found.
func getActiveEnvironment(workDir string) string {
	data, err := os.ReadFile(filepath.Join(workDir, ".env"))
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
