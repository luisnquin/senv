package main

import "fmt"

const DEFAULT_VERSION = "unversioned"

var (
	version = DEFAULT_VERSION
	commit  string
)

func main() {
	fmt.Printf("senv %s <%s>\n", version, commit)
}
