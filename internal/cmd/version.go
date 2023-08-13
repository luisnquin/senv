package cmd

import (
	"fmt"
	"os"

	"github.com/luisnquin/senv/internal"
)

func Version(version, commit string) error {
	if version == "" {
		version = internal.DEFAULT_VERSION
	}

	if commit != "" {
		commit = fmt.Sprintf("<%s>", commit)
	}

	_, err := fmt.Fprintf(os.Stdout, "%s %s %s\n", internal.PROGRAM_NAME, version, commit)

	return err
}
