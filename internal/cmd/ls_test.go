package cmd_test

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/luisnquin/senv/internal"
	"github.com/luisnquin/senv/internal/core"
	"github.com/luisnquin/senv/internal/test"
)

func TestLsCommandExpectedLines(t *testing.T) {
	settings := core.UserPreferences{
		Environments: []core.Environment{
			{
				Name: "dev",
				Variables: map[string]string{
					"development": "false",
				},
			},
			{
				Name: "prod",
				Variables: map[string]string{
					"development": "true",
				},
			},
		},
	}

	dir, err := test.Mkdir()
	must(t, err)

	err = test.WriteTo(dir, internal.DEFAULT_CFG_FILE_NAME, settings.Encode())
	must(t, err)

	must(t, os.Chdir(dir))

	code, out := test.RunMain("ls")
	if code != 0 {
		t.Errorf("unexpected exit code: %d", code)
		t.Fail()
	}

	linesMustContain := mapset.NewSet[string]()
	linesMustContain.Add(filepath.Join(dir, internal.DEFAULT_CFG_FILE_NAME))

	for _, env := range settings.Environments {
		linesMustContain.Add(env.Name)
	}

	s := bufio.NewScanner(strings.NewReader(out))

	for s.Scan() {
		line := strings.TrimPrefix(s.Text(), "- ")

		contains := false

		for _, subStr := range linesMustContain.ToSlice() {
			// Has scaped characters but still the original string can be recovered
			if strings.Contains(line, subStr) {
				linesMustContain.Remove(subStr)

				contains = true

				break
			}
		}

		if !contains {
			t.Errorf("ls output contains unexpected line: %s", line)
			t.Fail()
		}
	}
}

func must(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
