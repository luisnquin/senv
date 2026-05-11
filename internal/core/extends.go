package core

import (
	"fmt"

	"github.com/samber/lo"
)

// ResolveEnvironmentDefinition returns the named environment with the extends
// chain merged (same variable merge as when writing the .env file).
func (c *SenvConfig) ResolveEnvironmentDefinition(envName string) (EnvironmentDefinition, error) {
	environment, ok := lo.Find(c.Environments, func(e EnvironmentDefinition) bool {
		return e.Name == envName
	})
	if !ok {
		return EnvironmentDefinition{}, fmt.Errorf("environment not found")
	}

	if environment.Extends != "" {
		if err := mergeExtends(c, &environment); err != nil {
			return EnvironmentDefinition{}, err
		}
	}

	return environment, nil
}

func mergeExtends(preferences *SenvConfig, environment *EnvironmentDefinition) error {
	for _, other := range preferences.Environments {
		if other.Name != environment.Extends {
			continue
		}

		copyVars := lo.Assign(map[string]any{}, other.Variables)
		copyIgnoredCue := append([]string{}, other.IgnoredCueFiles...)
		copyCue := append([]CueDefinition{}, other.Cue...)

		for key, value := range environment.Variables {
			copyVars[key] = value
		}
		environment.Variables = copyVars
		environment.IgnoredCueFiles = append(copyIgnoredCue, environment.IgnoredCueFiles...)

		for _, originalDefinition := range environment.Cue {
			merged := false
			for index, extendedDefinition := range copyCue {
				if originalDefinition.File == extendedDefinition.File {
					mergedVars := lo.Assign(extendedDefinition.Variables, originalDefinition.Variables)
					copyCue[index] = CueDefinition{
						File:      originalDefinition.File,
						Variables: mergedVars,
					}
					merged = true
					break
				}
			}
			if !merged {
				copyCue = append(copyCue, originalDefinition)
			}
		}
		environment.Cue = copyCue

		return nil
	}

	return fmt.Errorf("environment '%s' could not be found (required by '%s')", environment.Extends, environment.Name)
}
