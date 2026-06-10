// main.go
package main

import "github.com/elricho/iicu/cmd"

// version is overridden at build time via -ldflags "-X main.version=...".
// GoReleaser and the Makefile both set it; defaults to "dev" for plain builds.
var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
