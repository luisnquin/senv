package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/luisnquin/senv/fsutils"
	"gopkg.in/yaml.v3"
)

var configFiles = []string{"senv.yaml", "senv.yml"}

type SenvConfig struct {
	// User defined environments.
	Environments []Environment  `yaml:"envs"`
	Defaults     map[string]any `yaml:"defaults"`
	// Options      map[string][]any `yaml:"options"`
}

func (c SenvConfig) validate() error {
	namesRegister := make(map[string]struct{}, len(c.Environments))

	for _, e := range c.Environments {
		if _, ok := namesRegister[e.Name]; ok {
			return fmt.Errorf("environment name %q already registered", e.Name)
		}

		namesRegister[e.Name] = struct{}{}
	}

	return nil
}

func loadConfig(workDirPath string) (*SenvConfig, error) {
	c := new(SenvConfig)

	if !fsutils.DirExists(workDirPath) {
		return nil, fmt.Errorf("provided work dir path does not exist")
	}

	workDirPath = filepath.Clean(workDirPath)

	for _, fileName := range configFiles {
		configPath := filepath.Join(workDirPath, fileName)

		info, err := os.Stat(configPath)
		if err != nil {
			continue
		}

		if !info.IsDir() {
			data, err := os.ReadFile(configPath)
			if err != nil {
				return nil, err
			}

			if err := yaml.Unmarshal(data, c); err != nil {
				return nil, err
			}

			if err := c.validate(); err != nil {
				return nil, err
			}

			return c, nil
		}
	}

	return nil, nil
}
