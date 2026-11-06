package commands

import (
	"fmt"
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/providers"
	"om/internal/ui"
)

func Install(args ...string) {
	if len(args) == 0 {
		logger.Error("no package provided")
		logger.Log("usage: om install [container] <pkg>")
		return
	}

	containerName := "arch" // TODO: infer containers or selection
	pkg := args[0]

	if len(args) > 1 {
		containerName = args[0]
		pkg = args[1]
	}

	container, err := containers.Find(containerName)
	if err != nil {
		logger.Errorf("%v\n", err)
		return
	}

	providerName, err := providers.Detect(container)
	if err != nil {
		logger.Errorf("%v\n", err)
		return
	}

	provider, err := providers.Get(providerName)
	if err != nil {
		logger.Errorf("%v\n", err)
		return
	}

	ui.PrintBanner()
	logger.Logf("installing \"%s\" using %s\n", pkg, provider.Name())
	fmt.Println()

	if err := provider.Install(container, pkg); err != nil {
		fmt.Println()
		logger.Errorf("failed to install package: %v\n", err)
		return
	}

	fmt.Println()
	logger.Successf("installed \"%s\"\n", pkg)
}
