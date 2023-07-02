package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/integrii/flaggy"
	"github.com/luisnquin/senv/env"
	"github.com/luisnquin/senv/log"
	"github.com/luisnquin/senv/prompt"
	"github.com/samber/lo"
)

const DEFAULT_VERSION = "unversioned"

var (
	version = DEFAULT_VERSION
	commit  string

	//go:embed help.tpl
	helpTpl string
	//go:embed docs/senv.example.yaml
	genericConfigFile []byte
)

func main() {
	check := flaggy.NewSubcommand("check")
	check.Description = "Checks wether the current working directory has `senv.yaml` or `.env` files"
	flaggy.AttachSubcommand(check, 1)

	to := flaggy.NewSubcommand("to")
	to.Description = "Allows you to switch to other environment without a prompt, directly as an argument"
	flaggy.AttachSubcommand(to, 1)

	ls := flaggy.NewSubcommand("ls")
	ls.Description = "List all the environments in the found working directory."
	flaggy.AttachSubcommand(ls, 1)

	revert := flaggy.NewSubcommand("revert")
	revert.Description = "Reverts the last environment switch"
	flaggy.AttachSubcommand(revert, 1)

	init := flaggy.NewSubcommand("init")
	init.Description = "Creates a new configuration file in the current directory"
	flaggy.AttachSubcommand(init, 1)

	flaggy.SetName("senv")
	flaggy.SetDescription("Switch between .env files")
	flaggy.SetVersion(fmt.Sprintf("senv %s <%s>", version, commit))
	flaggy.DefaultParser.SetHelpTemplate(helpTpl)
	flaggy.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	hasEnvOrEnvFiles := env.HasUsableWorkDir(currentDir)

	switch {
	case revert.Used:
		println("revert")
	case ls.Used:
		println("ls")
	case to.Used:
		println("to")
	case check.Used:
		if hasEnvOrEnvFiles {
			println("has YAML or env files")
		} else {
			println("doesn't have YAML or env files")
			os.Exit(1)
		}
	case init.Used:
		const configFilePath = "./senv.yaml"

		_, err := os.Stat(configFilePath)
		if err == nil {
			log.Pretty.Error("configuration file already exists in current directory")
		}

		err = os.WriteFile(configFilePath, genericConfigFile, os.ModePerm)
		if err != nil {
			log.Pretty.Fatal(err.Error())
		}
	default:
		if !hasEnvOrEnvFiles {
			log.Pretty.Error1("Current working folder doesn't have a `senv.yaml` or `.env` files")
		}

		workDir, err := env.ResolveUsableWorkDir(currentDir)
		if err != nil {
			log.Pretty.Fatal(err.Error())
		}

		environments, err := env.LoadEnvironments(workDir)
		if err != nil {
			log.Pretty.Fatal(err.Error())
		}

		envNames := make([]string, len(environments))

		for i, env := range environments {
			envNames[i] = env.Name
		}

		selected, ok := prompt.ListSelector("Select an environment", envNames)
		if !ok {
			os.Exit(1)
		}

		environment, ok := lo.Find(environments, func(e env.Environment) bool {
			return e.Name == selected
		})
		if !ok {
			log.Pretty.Fatal("lol")
		}

		dotEnv, err := env.GenerateDotEnv(environment)
		if err != nil {
			log.Pretty.Error(err.Error())
		}

		if err := os.WriteFile(filepath.Join(workDir, ".env"), dotEnv, os.ModePerm); err != nil {
			log.Pretty.Error(err.Error())
		}

		log.Pretty.Messagef("Switched to '%s'", environment.Name)
	}
}
