package containers

import (
	"fmt"
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
		return fmt.Errorf("container %q already exists", c.Name)
	}

	return c.create()
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

func (c Container) Run(args ...string) error {
	if err := c.Ensure(); err != nil {
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
