package providers

import (
	"om/internal/containers"
)

type Pacman struct{}

func (p Pacman) Name() string {
	return "pacman"
}

func (p Pacman) Install(container containers.Container, packages ...string) error {
	args := append([]string{
		"sudo", "pacman", "-S", "--noconfirm",
	}, packages...)

	return container.Run(args...)
}
