package containers

import (
	"fmt"
	"om/internal/lib/logger"
	"om/internal/ui"
	"os"
	"os/exec"
	"strings"
)

func (c Container) Exists() (bool, error) {
	cmd := exec.Command("distrobox-list")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	for _, line := range strings.Split(string(output), "\n") {
		columns := strings.Split(line, "|")

		if len(columns) < 2 {
			continue
		}

		name := strings.TrimSpace(columns[1])

		if name == c.Name {
			return true, nil
		}
	}

	return false, nil
}

func (c Container) Running() bool {
	cmd := exec.Command("podman", "inspect", "-f", "{{.State.Running}}", c.Name)
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}

func (c Container) create() error {
	cmd := exec.Command(
		"distrobox-create",
		"--name", c.Name,
		"--image", c.Image,
		"--yes",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"distrobox-create failed: %w\n%s",
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return nil
}

func (c Container) Create() error {
	exists, err := c.Exists()
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("container %q already exists\n", c.Name)
	}

	return c.create()
}

func (c Container) Remove() error {
	cmd := exec.Command("distrobox-rm", c.Name)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to remove container %q: %w\n%s",
			c.Name,
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return nil
}

func (c Container) Ensure() error {
	exists, err := c.Exists()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return c.create()
}

func (c Container) Init() error {
	ui.PrintBanner()
	logger.Logf("initialising container \"%q\"\n", c.Name)

	return c.InitSilently()
}

func (c Container) InitSilently() error {
	if err := c.StartSilently(); err != nil {
		return err
	}

	if err := c.StopSilently(); err != nil {
		return err
	}

	return nil
}

func (c Container) Start() error {
	ui.PrintBanner()
	logger.Logf("starting container %q\n", c.Name)
	return c.StartSilently()
}

func (c Container) StartSilently() error {
	cmd := exec.Command("podman", "start", c.Name)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to start container %q: %w\n%s",
			c.Name,
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return nil
}

func (c Container) Stop() error {
	ui.PrintBanner()
	logger.Logf("stopping container %q\n", c.Name)

	return c.StopSilently()
}

func (c Container) StopSilently() error {
	cmd := exec.Command("podman", "stop", c.Name)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to stop container %q: %w\n%s",
			c.Name,
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return nil
}

func (c Container) Run(args ...string) error {
	if err := c.Ensure(); err != nil {
		return err
	}

	if err := c.StartSilently(); err != nil {
		return err
	}

	cmd := exec.Command(
		"distrobox-enter",
		append([]string{c.Name, "--"}, args...)...,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (c Container) Output(args ...string) (string, error) {
	if err := c.Ensure(); err != nil {
		return "", err
	}

	cmd := exec.Command(
		"distrobox-enter",
		append([]string{c.Name, "--"}, args...)...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(
			"failed to execute command in container %q: %w\n%s",
			c.Name,
			err,
			strings.TrimSpace(string(output)),
		)
	}

	return string(output), nil
}

func Find(name string) (Container, error) {
	cmd := exec.Command("distrobox-list")
	output, err := cmd.Output()
	if err != nil {
		return Container{}, err
	}

	for _, line := range strings.Split(string(output), "\n") {
		columns := strings.Split(line, "|")

		if len(columns) < 4 {
			continue
		}

		containerName := strings.TrimSpace(columns[1])
		image := strings.TrimSpace(columns[3])

		if containerName == name {
			return New(containerName, image), nil
		}
	}

	return Container{}, fmt.Errorf("container %q not found", name)
}
