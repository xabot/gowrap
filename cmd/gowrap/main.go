package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/xabierlaiseca/gowrap/cmd/gowrap/commands"
)

func main() {
	parser := argparse.NewParser("gowrap", "gowrap util")
	versionsFileCommand, versionsFileCommandHandler := commands.NewVersionsFileCommand(&parser.Command)
	parser.Parse(os.Args)

	if versionsFileCommand.Happened() {
		exitOnError(versionsFileCommandHandler())
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
