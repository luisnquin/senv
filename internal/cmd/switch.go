package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"

	"cuelang.org/go/cue/cuecontext"
	cueformat "cuelang.org/go/cue/format"
	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/log"
	"github.com/luisnquin/senv/internal/prompt"
	"github.com/samber/lo"
)

// Creates a prompt selector that allows the user to select the environment to switch.
func Switch(currentDir string) error {
	if !core.HasConfigFiles(currentDir) {
		log.Pretty.Error1("Current working folder doesn't have a `senv.yaml`")
	}

	settings, err := core.LoadUserPreferences()
	if err != nil {
		return err
	}

	var activeEnv string

	if envFilePath, err := settings.GetEnvFilePath(); err == nil {
		activeEnv = getActiveEnvironment(envFilePath)
	}

	envNames := make([]string, len(settings.Environments))

	for i, env := range settings.Environments {
		envNames[i] = env.Name
	}

	selected, ok := prompt.ListSelector("Select an environment", envNames, activeEnv)
	if !ok {
		os.Exit(1)
	}

	if err := switchDotEnvFileFromName(settings, selected, settings.UseExportPrefix); err != nil {
		return err
	}

	log.Pretty.Messagef("Switched to '%s'", selected)

	return nil
}

func switchDotEnvFileFromName(preferences *core.SenvConfig, envToSwitch string, useExportPrefix bool) error {
	environment, ok := lo.Find(preferences.Environments, func(e core.EnvironmentDefinition) bool {
		return e.Name == envToSwitch
	})
	if !ok {
		return errors.New("environment not found")
	}

	groupedEnvVars := groupAndSortByPrefix(environment.Variables)

	dotEnvData, err := core.GenerateDotEnv(environment.Name, groupedEnvVars, useExportPrefix)
	if err != nil {
		return err
	}

	envFilePath, err := preferences.GetEnvFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(envFilePath, dotEnvData, os.ModePerm); err != nil {
		return err
	}

	if err := generateCueFiles(preferences, environment); err != nil {
		return err
	}

	return nil
}

func generateCueFiles(preferences *core.SenvConfig, environment core.EnvironmentDefinition) error {
	// global variables are overriden by local variables
	everyVariable := lo.Assign(preferences.GlobalVariables, environment.Variables)

	var err error

	// TODO: add tag if is in global scope
	writeCueFileFn := func(i int, cueVariables map[string]any, filePath string) error {
		ctx := cuecontext.New()
		value := ctx.Encode(cueVariables)

		cueData, err := cueformat.Node(value.Syntax())
		if err != nil {
			return fmt.Errorf("error formatting CUE file[%d] for environment %q: %w", i, environment.Name, err)
		}

		if err := os.WriteFile(filePath, cueData, os.ModePerm); err != nil {
			return fmt.Errorf("error writing CUE file[%d] for environment %q: %w", i, environment.Name, err)
		}

		return nil
	}

	if preferences.Cue != nil && len(environment.Cue) == 0 {
		for i, globalDefinition := range preferences.Cue.GlobalDefinitions {
			cueVariables, err := parseCueVarTplStrs(globalDefinition.Variables, everyVariable)
			if err != nil {
				return fmt.Errorf("error parsing CUE file[%d] variables for environment %q: %w,", i, environment.Name, err)
			}

			if err := writeCueFileFn(i, cueVariables, globalDefinition.File); err != nil {
				return err
			}
		}

		return nil
	}

	for i, cueDefinition := range environment.Cue {
		cueVariables := cueDefinition.Variables

		if preferences.Cue != nil {
			globalDefinition, found := lo.Find(preferences.Cue.GlobalDefinitions, func(d core.CueDefinition) bool {
				return d.File == cueDefinition.File
			})
			if found {
				// global variables are overriden by local variables
				cueVariables = lo.Assign(globalDefinition.Variables, cueVariables)
			}
		}

		cueVariables, err = parseCueVarTplStrs(cueVariables, everyVariable)
		if err != nil {
			return fmt.Errorf("error parsing CUE file[%d] variables for environment %q: %w,", i, environment.Name, err)
		}

		if err := writeCueFileFn(i, cueVariables, cueDefinition.File); err != nil {
			return err
		}
	}

	return nil
}

func parseCueVarTplStrs(cueVars, otherVars map[string]any) (map[string]any, error) {
	result := make(map[string]any, len(cueVars))

	for k, v := range cueVars {
		if s, ok := v.(string); ok {
			tpl, err := template.New("cue").Parse(s)
			if err != nil {
				return nil, err
			}

			var b bytes.Buffer
			if err = tpl.Execute(&b, map[string]any{
				"vars":      otherVars,
				"variables": otherVars,
			}); err != nil {
				return nil, err
			}

			result[k] = b.String()
		} else {
			result[k] = v
		}
	}

	return result, nil
}
