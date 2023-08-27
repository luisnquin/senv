package log

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

var Pretty = prettyPrinter{}

type prettyPrinter struct{}

func (p prettyPrinter) Message(s string) {
	fmt.Fprintf(os.Stdout, color.HEX("#74a7f2").Sprintf("%s\n", s))
}

func (p prettyPrinter) Messagef(template string, more ...any) {
	p.Message(fmt.Sprintf(template, more...))
}

func (p prettyPrinter) Warn(message string) {
	fmt.Fprint(os.Stderr, color.HEX("#ebe973").Sprintf("Warning: %s\n", message))
}

func (p prettyPrinter) Warnf(template string, more ...any) {
	p.Warn(fmt.Sprintf(template, more...))
}

func (p prettyPrinter) Error(message string) {
	fmt.Fprint(os.Stderr, color.HEX("#e63758").Sprintf("Error: %s\n", message))
}

func (p prettyPrinter) Error1(message string) {
	p.Error(message)
	os.Exit(1)
}

func (p prettyPrinter) Errorf(template string, more ...any) {
	p.Error(fmt.Sprintf(template, more...))
}

func (p prettyPrinter) Errorf1(template string, more ...any) {
	p.Errorf(template, more...)
	os.Exit(1)
}

func (p prettyPrinter) Fatal(message string) {
	fmt.Fprint(os.Stderr, color.HEX("#c41f3e").Sprintf("boom 💥, %s\n", message))
	os.Exit(1)
}

func (p prettyPrinter) Fatalf(template string, more ...any) {
	p.Fatal(fmt.Sprintf(template, more...))
}
