package senvtest

import (
	"bytes"
	"io"
	"os"

	"github.com/luisnquin/senv/internal"
	"github.com/luisnquin/senv/internal/app"
)

func RunMain(args ...string) (int, string) {
	os.Args = append([]string{"test"}, args...)

	r, w, _ := os.Pipe()

	oldOut := os.Stdout
	os.Stdout = w

	code := app.Run(internal.DEFAULT_VERSION, "")

	outC := make(chan string)

	go func() {
		var b bytes.Buffer

		io.Copy(&b, r)
		outC <- b.String()
	}()

	w.Close()
	os.Stdout = oldOut

	return code, <-outC
}
