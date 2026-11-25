package commands

import (
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/providers"
	"om/internal/ui"
)

func Doctor(args ...string) {
	if len(args) > 1 {
		logger.Error("too many arguments")
		logger.Log("usage: om doctor [container]")
		return
	}

	if len(args) == 0 { // todo: doctor all containers
		logger.Error("please specify a container")
		logger.Log("usage: om doctor <container>")
		return
	}

	ui.PrintBanner()
	logger.Logf("diagnosing container %q\n\n", args[0])

	container, err := containers.Find(args[0])
	if err != nil {
		logger.Errorf("failed to find the container, are you sure it exists?\n")
		return
	}

	// wasRunning, err := container.Running()
	status, err := container.Status()
	if err != nil {
		logger.Errorf("failed to check container status: %v\n", err)
		return
	}

	if err != nil {
		logger.Errorf("failed to check container status: %v\n", err)
		return
	}

	logger.Log("container found, checking if it is running...")
	if !status.Running {
		logger.Log("container is not running, starting it...")
		if err := container.StartSilently(); err != nil {
			logger.Errorf("failed to start the container: %v\n", err)
			return
		}
	}

	if !status.Running {
		defer func() {
			logger.Log("restoring original container state...")

			if err := container.StopSilently(); err != nil {
				logger.Errorf("failed to restore container state: %v\n", err)
			}

			logger.Success("diagnosis complete")
		}()
	}

	logger.Log("container is running, proceeding with diagnosis...\n")

	_, err = container.Output("true")
	if err != nil {
		logger.Errorf("container test execution failed: %v\n", err)
		logger.Warn("container might need further investigation")
		return
	}

	logger.Successf("container test execution succeeded\n")

	logger.Log("detecting package manager...")
	provider, err := providers.Find(container)
	if err != nil {
		logger.Errorf("package manager detection failed: %v\n", err)
		return
	}

	logger.Successf("detected package manager: %s\n\n", provider.Name())

	if status.Running {
		logger.Success("diagnosis complete")
	}
}
