package commands

import (
	"fmt"
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/ui"
	"om/internal/ui/colours"

	"github.com/fatih/color"
)

func List(args ...string) {
	if len(args) != 0 {
		logger.Error("too many arguments")
		logger.Log("usage: om list")
		return
	}

	ui.PrintBanner()

	logger.Logf("finding containers...")

	cs, err := containers.List()
	if err != nil {
		logger.Errorf("failed to list containers: %v\n", err)
		return
	}

	statuses, err := containers.Statuses(cs)
	if err != nil {
		logger.Errorf("failed to get container statuses: %v\n", err)
		return
	}

	fmt.Printf("\x1b[2K\r")

	if len(cs) == 0 {
		logger.Log("no containers found")
		return
	}

	logger.Log("containers:")

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

		fmt.Printf("%s %s", c.Name, colours.Builder("("+c.Image+") - "+status, color.Faint))
		fmt.Println()
	}
}
