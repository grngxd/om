package commands

import (
	"om/internal/containers"
	"om/internal/lib/logger"
)

func Run(args ...string) {
	if len(args) == 0 {
		logger.Error("no command provided")
		return
	}

	// container := containers.Arch{}
	containerName := "arch"
	if len(args) > 1 {
		containerName = args[0]
		args = args[1:]
	}

	var container containers.Container

	switch containerName {
	case "arch":
		container = containers.Arch{}
	default:
		logger.Errorf("unsupported container: %s\n", containerName)
		return
	}

	if err := container.Run(args...); err != nil {
		logger.Errorf("failed to run command: %v\n", err)
	}
}
