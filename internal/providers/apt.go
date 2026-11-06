package providers

import (
	"om/internal/containers"
)

type Apt struct{}

func (a Apt) Name() string {
	return "apt-get"
}

func (a Apt) Install(container containers.Container, packages ...string) error {
	if err := container.Run("sudo", "apt-get", "update", "-qq"); err != nil {
		return err
	}

	args := append([]string{"sudo", "apt-get", "install", "-y"}, packages...)

	return container.Run(args...)
}
