package main

import (
	"github.com/Nchezhegova/gophkeeper/cmd/client/commands"
)

var (
	version   string
	buildDate string
)

func main() {
	commands.Version = version
	commands.BuildDate = buildDate

	commands.Execute()
}
