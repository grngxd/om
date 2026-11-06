package ui

import (
	"fmt"
	"om/internal/constants"
	"strings"
)

func PrintBanner() {
	avatar := strings.Trim(`
██  ██  ██
██████████
  ██  ██
██      ██
██      ██`, "\n")

	fmt.Printf("\n%s om %s\n\n", avatar, constants.Version)
}
