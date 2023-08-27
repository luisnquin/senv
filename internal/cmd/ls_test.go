package cmd_test

import (
	"testing"

	"github.com/luisnquin/senv/internal/test"
)

func TestLs(t *testing.T) {
	code, out := test.RunMain("ls")

	t.Log(code, out)
}
