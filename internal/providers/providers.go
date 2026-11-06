package providers

import (
	"fmt"
	"om/internal/containers"
)

type Provider interface {
	Name() string
	Install(container containers.Container, packages ...string) error
}

func Get(provider string) (Provider, error) {
	switch provider {
	case "pacman":
		return Pacman{}, nil
	case "apt-get":
		return Apt{}, nil
	case "dnf":
		return Dnf{}, nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}

func Detect(c containers.Container) (string, error) {
	providers := []string{
		"pacman",
		"apt-get",
		"dnf",
	}

	for _, provider := range providers {
		_, err := c.Output("sh", "-c", "command -v "+provider)
		if err == nil {
			return provider, nil
		}
	}

	return "", fmt.Errorf("no supported package manager found, this should not happen")
}
