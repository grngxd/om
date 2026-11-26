package commands

import (
	"fmt"
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/ui/colours"

	"github.com/fatih/color"
)

func Run(args ...string) {
	if len(args) == 0 {
		logger.Error("no command provided")
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
