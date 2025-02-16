package core

import (
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
	EnvFile string `yaml:"env_file"`
	// Indicates whether to use the 'export' prefix in the final .env file or not.
	UseExportPrefix bool `yaml:"use_export_prefix"`
	// The working directory absolute path.
	WorkDirectory string `yaml:"-"`
	// The path of the file associated with the loaded preferences.
	SourceFilePath string `yaml:"-"`
}

// An user defined environment.
type Environment struct {
	// The environment name.
	Name string `yaml:"name"`
	// Key value pairs for the environment.
	//
	// Is expected when serialized it will look
	// like this:
	//
	// 	FOO=bar
	// 	BAR=foo
	Variables map[string]any `yaml:"variables"`
}

// Find, validate and deserialize the senv.yaml or senv.yml files in the
// current directory or parents.
func LoadUserPreferences() (*UserPreferences, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	workDirPath := getDirWithConfig(currentPath)

	c := &UserPreferences{
		WorkDirectory: workDirPath,
	}

	for _, fileName := range getConfigFiles() {
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

	return nil, getErrConfigNotFound()
}

// Returns the destiny path of the .env file to be generated.
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
