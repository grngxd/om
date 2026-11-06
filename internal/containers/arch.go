package containers

import "om/internal/distrobox"

type Arch struct{}

func (a Arch) Name() string {
	return "arch"
}

func (a Arch) Ensure() error {
	return distrobox.Ensure(a.Name(), "archlinux:latest")
}

func (a Arch) Run(args ...string) error {
	if err := a.Ensure(); err != nil {
		return err
	}

	return distrobox.Enter(a.Name(), args...)
}

func (a Arch) Install(pkg string) error {
	if err := a.Ensure(); err != nil {
		return err
	}

	return distrobox.Enter(
		a.Name(),
		"sudo", "pacman", "-S", "--noconfirm", pkg,
	)
}
