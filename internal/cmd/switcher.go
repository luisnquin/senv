package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"os"

	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/log"
	"github.com/samber/lo"
)

type Switcher struct{ settings *core.UserPreferences }

func NewSwitcher(settings *core.UserPreferences) Switcher { return Switcher{settings} }

func (s Switcher) GetActiveEnvironment() (string, error) {
	path, err := s.settings.GetEnvFilePath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line := scanner.Text()

		if len(rxSenvDotEnvComment.FindAllString(line, -1)) == 2 {
			return rxSenvDotEnvComment.ReplaceAllString(line, ""), nil
		}
	}

	return "", nil
}

func (s Switcher) Switch(name string) error {
	findEnvironmentFn := func(e core.Environment) bool { return e.Name == name }

	environment, ok := lo.Find(s.settings.Environments, findEnvironmentFn)
	if !ok {
		return errors.New("environment not found")
	}

	binds := make(map[string]string, len(s.settings.Binds))

	for name, b := range s.settings.Binds {
		if b.Use == "" && len(b.Options) > 0 {
			binds[name] = b.Options[0]
		} else {
			if !lo.Contains(b.Options, b.Use) {
				log.Pretty.Warnf("'%s' is not between 'binds.%s.options'", b.Use, name)
			}

			binds[name] = b.Use
		}
	}

	dotEnvData, err := core.GenerateDotEnv(environment, binds, s.settings.UseExportPrefix)
	if err != nil {
		return err
	}

	envFilePath, err := s.settings.GetEnvFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(envFilePath, dotEnvData, os.ModePerm); err != nil {
		return err
	}

	return nil
}
