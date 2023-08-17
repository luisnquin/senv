package core

import "errors"

func getErrUnableToResolveWorkDir() error {
	return errors.New("work directory cannot be resolved")
}

func getErrConfigNotFound() error {
	return errors.New("senv.yaml file not found")
}
