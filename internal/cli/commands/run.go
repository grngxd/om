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

	containerName := "arch" // TODO: actual real inference

	if len(args) > 1 {
		containerName = args[0]
		args = args[1:]
	}

	container, err := containers.Find(containerName)
	if err != nil {
		logger.Errorf("%v\n", err)
		return
	}

	if err := container.Run(args...); err != nil {
		logger.Errorf("failed to run command: %v\n", err)
	}
}
