package providers

import "om/internal/containers"

type Dnf struct{}

func (d Dnf) Name() string {
	return "dnf"
}

func (d Dnf) Install(container containers.Container, packages ...string) error {
	args := append([]string{"sudo", "dnf", "install", "-y"}, packages...)

	return container.Run(args...)
}
