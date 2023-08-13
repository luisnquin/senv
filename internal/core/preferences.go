package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// The user preferences.
type UserPreferences struct {
	// User defined environments.
	Environments []Environment `yaml:"envs"`
	// Default variables for all the environments.
	Defaults map[string]any `yaml:"defaults"`
	// The relative or absolute path to the env file.
	EnvFile string `yaml:"envFile"`
	// Indicates whether to use the 'export' prefix in the final .env file or not.
	UseExportPrefix bool `yaml:"useExportPrefix"`
	// The working directory absolute path.
	WorkDirectory string `yaml:"-"`
	// The path of the file associated with the loaded preferences.
	SourceFilePath string `yaml:"-"`
}

// An user defined environment.
type Environment struct {
	Name      string         `yaml:"name"`
	Variables map[string]any `yaml:"variables"`
}

var ErrConfigNotFound = errors.New("senv.yaml file not found")

func LoadUserPreferences() (*UserPreferences, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	workDirPath := workDirHasProgramFiles(currentPath, false)

	c := &UserPreferences{
		WorkDirectory: workDirPath,
	}

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

			c.SourceFilePath = configPath

			return c, nil
		}
	}

	return nil, ErrConfigNotFound
}

func (c UserPreferences) GetEnvFilePath() (string, error) {
	if filepath.IsAbs(c.EnvFile) {
		return c.EnvFile, nil
	}

	return filepath.Join(c.WorkDirectory, c.EnvFile), nil
}

func (c *UserPreferences) validate() error {
	namesRegister := make(map[string]struct{}, len(c.Environments))

	for _, e := range c.Environments {
		if _, ok := namesRegister[e.Name]; ok {
			return fmt.Errorf("environment name %q already registered", e.Name)
		}

		for k, v := range c.Defaults {
			_, ok := e.Variables[k]
			if !ok {
				e.Variables[k] = v
			}
		}

		namesRegister[e.Name] = struct{}{}
	}

	if c.EnvFile == "" {
		c.EnvFile = "./.env"
	}

	info, err := os.Stat(c.EnvFile)
	if err == nil && info.IsDir() {
		return fmt.Errorf("specified env file is a folder: %s", c.EnvFile)
	}

	return nil
}
