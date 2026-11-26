package containers

import (
	"fmt"
	"io"
	"om/internal/lib/logger"
	"om/internal/ui"
	"os"
	"os/exec"
	"strings"
)

func List() ([]Container, error) {
	cmd := exec.Command("distrobox-list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var containers []Container

	for _, line := range strings.Split(string(output), "\n")[1:] {
		columns := strings.Split(line, "|")

		if len(columns) < 4 {
			continue
		}

		name := strings.TrimSpace(columns[1])
		image := strings.TrimSpace(columns[3])

		container := New(name, image)
		if err := container.Ensure(); err != nil {
			return nil, fmt.Errorf("container does not exist %q: %w", name, err)
		}

		containers = append(containers, container)
	}

	return containers, nil
}

func Statuses(containers []Container) (map[string]string, error) {
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

type ContainerStatus struct {
	Running    bool
	Exited     bool
	Created    bool
	Configured bool
	Paused     bool
}

func (c Container) Status() (ContainerStatus, error) {
	cmd := exec.Command("podman", "inspect", "-f", "{{.State.Status}}", c.Name)
	output, err := cmd.Output()
	if err != nil {
		return ContainerStatus{}, err
	}

	src := strings.TrimSpace(string(output))
	status := ContainerStatus{}

	switch src {
	case "running":
		status.Running = true
	case "exited":
		status.Exited = true
	case "created":
		status.Created = true
	case "configured":
		status.Configured = true
	case "paused":
		status.Paused = true
	default:
		logger.Warnf("unknown container status: %s\n", src)
		return ContainerStatus{}, fmt.Errorf("unknown container status: %s", src)
	}

	statuses := []bool{status.Running, status.Exited, status.Created, status.Configured, status.Paused}
	count := 0
	for _, s := range statuses {
		if s {
			count++
		}
	}

	if count != 1 {
		logger.Errorf("container %q has multiple/no statuses set, this should not happen\n", c.Name)
	}

	return status, nil
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

	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	// _, err := c.Output(
	// 	"sh", "-c",
	// 	`command -v "$1" >/dev/null 2>&1`,
	// 	"om", args[0],
	// )

	// faster to check if the command exists in the container by running it directly and checking for an error
	_, err := c.Output(
		"sh", "-c",
		`"$@"`,
		"om", args[0],
	)

	if err != nil {
		if strings.Contains(err.Error(), "command not found") {
			return fmt.Errorf("command %q not found in container, check the command name or install the package", args[0])
		}

		return fmt.Errorf("failed to execute command in container: %w", err)
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

	var stderr strings.Builder
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	err = cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf(
				"command failed with exit code %d: %s",
				exitErr.ExitCode(),
				strings.TrimSpace(stderr.String()),
			)
		}

		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
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
