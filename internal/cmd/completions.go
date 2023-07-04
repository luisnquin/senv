package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func Completion(completionsFolder fs.FS, shell string) error {
	compFilePath := fmt.Sprintf("completions/completions.%s", shell)

	f, err := completionsFolder.Open(compFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("completion script not found")
		}

		return err
	}

	defer f.Close()

	_, err = io.Copy(os.Stdout, f)

	return err
}
