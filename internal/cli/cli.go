package cli

import (
	"om/internal/cli/commands"
	"os"
)

func Execute() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		commands.Usage(false, os.Args[1:]...)
		return
	}

	switch os.Args[1] {
	case "install", "i":
		commands.Install(os.Args[2:]...)
	case "run":
		commands.Run(os.Args[2:]...)
	case "version", "v":
		commands.Version()
	case "create":
		commands.Create(os.Args[2:]...)
	default:
		commands.Usage(true, os.Args[1:]...)
	}
}
