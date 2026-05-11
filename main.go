package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/luisnquin/flaggy"
	"github.com/luisnquin/senv/internal"
	"github.com/luisnquin/senv/internal/assets"
	"github.com/luisnquin/senv/internal/cmd"
	"github.com/luisnquin/senv/internal/log"
)

var (
	version = internal.DEFAULT_VERSION
	commit  string
)

type flags struct {
	Get             bool
	Set             string
	CompletionShell string
	Ls              struct {
		Raw bool
	}
	Export struct {
		Format string
		Dir    string
		Pretty bool
		Env    string
	}
}

func main() {
	var flags flags

	flaggy.String(&flags.Set, "", "set", "Set the current environment skipping the default prompt")
	flaggy.Bool(&flags.Get, "", "get", "Get the current environment")

	ls := flaggy.NewSubcommand("ls")
	ls.Description = "List all the environments in the working directory"
	ls.Bool(&flags.Ls.Raw, "", "raw", "list environment without especial decorations")
	flaggy.AttachSubcommand(ls, 1)

	init := flaggy.NewSubcommand("init")
	init.Description = "Creates a new configuration file in the current directory"
	flaggy.AttachSubcommand(init, 1)

	export := flaggy.NewSubcommand("export")
	export.Description = "Print resolved variables for an environment (dotenv or json) to stdout"
	export.String(&flags.Export.Format, "f", "format", "Output format: dotenv or json")
	export.String(&flags.Export.Dir, "d", "dir", "Directory to resolve senv.yaml from (default: cwd)")
	export.Bool(&flags.Export.Pretty, "", "pretty", "Pretty-print JSON (json format only)")
	export.AddPositionalValue(&flags.Export.Env, "environment", 1, true, "Environment name from senv.yaml")
	flaggy.AttachSubcommand(export, 1)

	var completionShellArg string

	completion := flaggy.NewSubcommand("completion")
	completion.Hidden = true
	flaggy.AttachSubcommand(completion, 1)
	completion.AddPositionalValue(&completionShellArg, "shell", 1, true, "Supported shells: zsh && bash")

	flaggy.SetName(internal.PROGRAM_NAME)
	flaggy.SetDescription("Switch your .env file")
	flaggy.SetVersion(getVersion())
	flaggy.DefaultParser.SetHelpTemplate(assets.GetHelpTpl())

	code, err := flaggy.Parse()
	if err != nil {
		// log.Pretty.Error(err.Error())

		flaggy.ShowHelp(err.Error())
		os.Exit(code)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch {
	case completion.Used:
		if err := cmd.Completion(completionShellArg); err != nil {
			log.Pretty.Error(err.Error())
		}
	case ls.Used:
		if err := cmd.Ls(currentDir, flags.Ls.Raw); err != nil {
			log.Pretty.Error(err.Error())
		}
	case flags.Get:
		if err := cmd.GetEnv(); err != nil {
			log.Pretty.Fatal(err.Error())
		}
	case flags.Set != "":
		if err := cmd.SetEnv(currentDir, flags.Set); err != nil {
			log.Pretty.Error(err.Error())
		}
	case init.Used:
		if err := cmd.Init(); err != nil {
			log.Pretty.Error(err.Error())
		}
	case export.Used:
		if err := cmd.Export(currentDir, flags.Export.Dir, flags.Export.Env, flags.Export.Format, flags.Export.Pretty); err != nil {
			log.Pretty.Error(err.Error())
		}
	default:
		if err := cmd.Switch(currentDir); err != nil {
			log.Pretty.Error(err.Error())
		}
	}
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
