package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/luisnquin/senv/internal/core"
)

// Export prints resolved variables for the named environment to stdout. workDir
// is the directory used to locate senv.yaml (typically the process cwd). If
// configDir is non-empty, it is used instead of workDir for config resolution.
func Export(workDir, configDir, envName, format string, prettyJSON bool) error {
	envName = strings.TrimSpace(envName)
	if envName == "" {
		return errors.New("environment name is required")
	}

	resolveFrom := workDir
	if strings.TrimSpace(configDir) != "" {
		resolveFrom = configDir
	}

	if !core.HasConfigFiles(resolveFrom) {
		return errors.New("current working folder doesn't have a `senv.yaml`")
	}

	settings, err := core.LoadUserPreferencesFromDir(resolveFrom)
	if err != nil {
		return err
	}

	environment, err := settings.ResolveEnvironmentDefinition(envName)
	if err != nil {
		return err
	}

	f := strings.ToLower(strings.TrimSpace(format))
	if f == "" {
		f = exportFormatDotenv
	}

	switch f {
	case exportFormatDotenv:
		grouped := groupAndSortByPrefix(environment.Variables)
		out, err := core.GenerateDotEnv(environment.Name, grouped, settings.UseExportPrefix, false)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(out)

		return err

	case exportFormatJSON:
		grouped := groupAndSortByPrefix(environment.Variables)
		flat := flattenGroupedVariables(grouped)
		var payload []byte
		if prettyJSON {
			payload, err = json.MarshalIndent(flat, "", "  ")
		} else {
			payload, err = json.Marshal(flat)
		}
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(os.Stdout, string(payload))

		return err

	default:
		return fmt.Errorf("unknown format %q (use %s or %s)", format, exportFormatDotenv, exportFormatJSON)
	}
}

func flattenGroupedVariables(grouped []map[string]any) map[string]any {
	out := make(map[string]any)
	for _, block := range grouped {
		for k, v := range block {
			out[k] = v
		}
	}

	return out
}
