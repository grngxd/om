package commands

import (
	"fmt"
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/providers"
	"om/internal/ui"
	"om/internal/ui/colours"

	"github.com/fatih/color"
)

func Install(args ...string) {
	if len(args) == 0 {
		logger.Error("no package provided")
		logger.Log("usage: om install [container] <pkg>")
		return
	}

	// containerName := "arch" // TODO: actual real inference
	var containerName string
	logger.Logf("fetching containers...")

	cs, err := containers.List()
	if err != nil {
		logger.Errorf("failed to list containers: %v\n", err)
		return
	}

	if len(cs) == 0 {
		logger.Error("no containers found")
		return
	} else if len(cs) == 1 {
		containerName = cs[0].Name
	} else {
		statuses, err := containers.Statuses(cs)
		if err != nil {
			logger.Errorf("failed to get container statuses: %v\n", err)
			return
		}

		fmt.Printf("\x1b[2K\r")
		logger.Log("choose a container:")

		for _, c := range cs {
			fmt.Printf("  ")
			status := statuses[c.Name]

			switch status {
			case "running":
				fmt.Printf(colours.Builder("● ", color.FgGreen))
			case "exited":
				fmt.Printf(colours.Builder("○ ", color.FgRed))
			case "paused":
				fmt.Printf(colours.Builder("○ ", color.FgYellow))
			default:
				fmt.Printf(colours.Builder("○ ", color.Faint))
			}

			fmt.Printf("%s %s", c.Name, colours.Builder("("+c.Image+")", color.Faint))
			fmt.Println()
		}

		fmt.Printf(colours.Builder("  > ", color.Faint))
		fmt.Scanln(&containerName)
	}

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

	status, err := container.Status()
	if err != nil {
		logger.Errorf("failed to get container status: %v\n", err)
		return
	}

	if !status.Running {
		if err := container.Start(); err != nil {
			logger.Errorf("failed to start container: %v\n", err)
			return
		}
	}

	if !status.Running {
		defer func() {
			if err := container.Stop(); err != nil {
				logger.Errorf("failed to stop container: %v\n", err)
			}
		}()
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
