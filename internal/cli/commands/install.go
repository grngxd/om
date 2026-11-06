package commands

import (
	"om/internal/containers"
	"om/internal/logger"
	"om/internal/ui"
)

func Install(args ...string) {
	if len(args) == 0 {
		logger.Error("no package provided")
		logger.Log("usage: om install [distro?] <pkg>")
		return
	}

	pkg := args[0]
	containerName := "arch"

	// om i arch nvim
	if len(args) == 2 {
		containerName = args[0]
		pkg = args[1]
	}

	var container containers.Container

	switch containerName {
	case "arch":
		container = containers.Arch{}
	default:
		logger.Errorf("unrecognised container: %s\n", containerName)
		return
	}

	ui.PrintBanner()

	logger.Logf("installing \"%s\"\n", pkg)

	if err := container.Install(pkg); err != nil {
		logger.Errorf("failed to install package: %v\n", err)
		return
	}

	logger.Logf("installed \"%s\"\n", pkg)
}
