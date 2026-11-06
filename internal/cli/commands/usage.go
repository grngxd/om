package commands

import (
	"fmt"
	"om/internal/logger"
	"om/internal/ui"
)

func Usage(unrecognisedCommand bool, args ...string) {
	ui.PrintBanner()

	if unrecognisedCommand {
		logger.Logf("unrecognised command: \"%s\"\n", args[0])
	}

	logger.Log("usage:")
	fmt.Println(" om install [distro?] <pkg>")
	fmt.Println(" om run [distro?] <cmd>")
	fmt.Println(" om version")
}
