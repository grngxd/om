package runner

import (
	"om/internal/lib/logger"
	"os"
	"os/exec"
)

func Run(cmd string, args ...string) error {
	exe := exec.Command(cmd, args...)

	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	exe.Stdin = os.Stdin

	err := exe.Run()
	if err != nil {
		logger.Errorf("command failed: %v\n", err)
		return err
	}

	return nil
}
