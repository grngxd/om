package distrobox

import (
	"bytes"
	"os/exec"

	"om/internal/runner"
)

func Enter(container string, args ...string) error {
	return runner.Run(
		"distrobox-enter",
		append([]string{container, "--"}, args...)...,
	)
}

func Create(container string, image string) error {
	return runner.Run(
		"distrobox-create",
		"--name", container,
		"--image", image,
	)
}

func Exists(container string) bool {
	cmd := exec.Command("distrobox-list")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return false
	}

	return bytes.Contains(out.Bytes(), []byte(container))
}

func Ensure(container string, image string) error {
	if Exists(container) {
		return nil
	}
	return Create(container, image)
}
