package cli

import (
	"om/internal/cli/commands"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		commands.Usage()
		return
	}

	switch os.Args[1] {
	case "install", "i":
		commands.Install(os.Args[2:]...)
	case "run":
		commands.Run(os.Args[2:]...)
	case "version", "v":
		commands.Version()
	default:
		commands.Usage()
	}
}
