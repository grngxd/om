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
	// misc
	case "version", "v":
		commands.Version()

	// containers
	case "list", "ls":
		commands.List(os.Args[2:]...)
	case "create", "c":
		commands.Create(os.Args[2:]...)
	case "remove", "rm":
		commands.Remove(os.Args[2:]...)
	case "doctor", "doc":
		commands.Doctor(os.Args[2:]...)

	// packages
	case "install", "i":
		commands.Install(os.Args[2:]...)
	// case "uninstall", "un":
	// 	commands.Uninstall(os.Args[2:]...)

	case "run":
		commands.Run(os.Args[2:]...)
	default:
		if len(os.Args) == 1 {
			commands.Usage(true, os.Args[1:]...)
		}

		commands.Run(os.Args[1:]...)
	}
}
