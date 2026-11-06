package ui

import (
	"fmt"
	"om/internal/constants"
	"om/internal/lib/colours"
	"strings"
)

func PrintBanner() {
	avatar := strings.Trim(`
██  ██  ██
██████████
  ██  ██
██      ██
██  ██  ██
`, "\n")

	// fmt.Printf("\n%s om %s\n\n", colours.Green(avatar), strings.Split(constants.Version, " ")[0]+" "+colours.Green(strings.Split(constants.Version, " ")[1]))
	// fmt.Printf("\n%s om %s %s %s\n\n", colours.Green(avatar), constants.SemanticVersion, colours.Faint("+"), colours.Green(constants.BuildNumber))
	// fmt.Printf("\n%s om %s %s\n\n", colours.Green(avatar), constants.SemanticVersion, colours.Green(constants.BuildNumber))
	fmt.Printf("\n%s %s %s %s\n\n", colours.Green(avatar), colours.Faint("om"), constants.SemanticVersion, colours.Green(constants.BuildNumber))
}
