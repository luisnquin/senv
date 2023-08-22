package cmd

import (
	"bufio"
	"bytes"
	"os"
)

// Returns an empty string if not found.
func getActiveEnvironment(envFilePath string) string {
	data, err := os.ReadFile(envFilePath)
	if err != nil {
		return ""
	}

	s := bufio.NewScanner(bytes.NewReader(data))

	for s.Scan() {
		line := s.Text()

		if len(rxSenvDotEnvComment.FindAllString(line, -1)) == 2 {
			return rxSenvDotEnvComment.ReplaceAllString(line, "")
		}
	}

	return ""
}
