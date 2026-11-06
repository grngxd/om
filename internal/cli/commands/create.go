package commands

import (
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/ui"
)

func Create(args ...string) {
	if len(args) != 2 {
		logger.Error("invalid arguments")
		logger.Log("usage: om create <name> <image>")
		return
	}

	name := args[0]
	image := args[1]

	container := containers.New(name, image)

	ui.PrintBanner()
	logger.Logf("creating container \"%s\" with image \"%s\"\n", name, image)

	if err := container.Create(); err != nil {
		logger.Errorf("failed to create container: %v\n", err)
		return
	}

	logger.Successf("created container \"%s\"\n", name)
}
