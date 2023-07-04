package env

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
	// The working directory absolute path.
	WorkDirectory string `yaml:"-"`
	// The path of the file associated with the loaded preferences.
	SourceFilePath string `yaml:"-"`
}

var configFiles = []string{"senv.yaml", "senv.yml"}

var ErrConfigNotFound = errors.New("configuration file not found")

// Options      map[string][]any `yaml:"options"`

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

func LoadUserPreferences() (*UserPreferences, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	workDirPath := resolveUsableWorkDirectory(currentPath, false)

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
