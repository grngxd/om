package commands

import (
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/ui"
)

// om remove <container>
func Remove(args ...string) {
	if len(args) != 1 {
		logger.Error("invalid arguments")
		logger.Log("usage: om remove <name>")
		return
	}

	name := args[0]

	ui.PrintBanner()
	logger.Logf("removing container \"%s\"", name)

	container, err := containers.Find(name)
	if err != nil {
		logger.Errorf("failed to find container: %v\n", err)
		return
	}

	if err := container.Remove(); err != nil {
		logger.Errorf("failed to remove container: %v\n", err)
		return
	}
}
