package commands

import (
	"fmt"
	"om/internal/containers"
	"om/internal/lib/logger"
	"om/internal/ui"
	"om/internal/ui/colours"
	"os/exec"
	"strings"

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

	containers, err := containers.List()
	if err != nil {
		logger.Errorf("failed to list containers: %v\n", err)
		return
	}

	statuses, err := Statuses(containers)
	if err != nil {
		logger.Errorf("failed to get container statuses: %v\n", err)
		return
	}

	fmt.Printf("\x1b[2K\r")

	if len(containers) == 0 {
		logger.Log("no containers found")
		return
	}

	logger.Log("containers:")

	for _, c := range containers {
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

func Statuses(containers []containers.Container) (map[string]string, error) {
	if len(containers) == 0 {
		return map[string]string{}, nil
	}

	args := []string{
		"inspect",
		"--format",
		"{{.Name}}|{{.State.Status}}",
	}

	for _, c := range containers {
		args = append(args, c.Name)
	}

	cmd := exec.Command("podman", args...)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	statuses := make(map[string]string)

	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		columns := strings.Split(line, "|")
		if len(columns) != 2 {
			continue
		}

		name := strings.TrimPrefix(strings.TrimSpace(columns[0]), "/")
		status := strings.TrimSpace(columns[1])

		statuses[name] = status
	}

	return statuses, nil
}
