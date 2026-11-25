package commands

import (
	"fmt"
	"om/internal/lib/logger"
	"om/internal/ui"
)

func Usage(unrecognisedCommand bool, args ...string) {
	ui.PrintBanner()

	if unrecognisedCommand {
		logger.Logf("unrecognised command: \"%s\"\n", args[0])
	}

	logger.Log("usage:\n")
	fmt.Println(" containers:")

	fmt.Println("   om list")
	fmt.Println("   om create <name> <image>")
	fmt.Println("   om remove <name>")
	fmt.Println("   om doctor <name>\n")

	fmt.Println(" packages:")
	fmt.Println("   om install [distro?] <package>")
	fmt.Println("   om run [distro?] <command>\n")

	fmt.Println(" miscellaneous:")
	fmt.Println("   om version")
}
