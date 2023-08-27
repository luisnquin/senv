package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/luisnquin/senv/internal"
	"github.com/luisnquin/senv/internal/app"
)

var (
	version = internal.DEFAULT_VERSION
	commit  string
)

func main() {
	os.Exit(app.Run(getProgramVersion()))
}

func getCommit() string {
	if commit != "" {
		return commit
	}

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, kv := range info.Settings {
			if kv.Key == "vcs.revision" {
				return kv.Value
			}
		}
	}

	return ""
}

func getProgramVersion() string {
	if version == "" {
		version = internal.DEFAULT_VERSION
	}

	commit := getCommit()
	if commit != "" {
		commit = fmt.Sprintf("<%s>", commit)
	}

	return fmt.Sprintf("%s %s %s", internal.PROGRAM_NAME, version, commit)
}
