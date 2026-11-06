package ui

import (
	"fmt"
	"om/internal/constants"
	"om/internal/ui/colours"
	"strings"

	"github.com/fatih/color"
)

func PrintBanner() {
	avatar := strings.Trim(`
██  ██  ██
██████████
  ██  ██
██      ██
██  ██  ██
`, "\n")

	fmt.Printf("\n%s %s %s %s\n\n", colours.Green(avatar), colours.Faint("om"), constants.SemanticVersion, colours.Builder(constants.BuildNumber, color.FgGreen, color.Bold))
}
