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
	os.Exit(app.Run(getVersion(), getCommit()))
}

func getCommit() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, kv := range info.Settings {
			if kv.Key == "vcs.revision" {
				return kv.Value
			}
		}
	}

	return commit
}

func getVersion() string {
	if version == "" {
		version = internal.DEFAULT_VERSION
	}

	if commit != "" {
		commit = fmt.Sprintf("<%s>", commit)
	}

	return fmt.Sprintf("%s %s %s", internal.PROGRAM_NAME, version, commit)
}
