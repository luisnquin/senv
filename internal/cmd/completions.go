package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/luisnquin/senv/internal/assets"
)

func Completion(shell string) error {
	completionsFolder := assets.GetCompletionsFolder()
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
