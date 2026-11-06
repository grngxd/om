package commands

import (
	"fmt"
	"om/internal/constants"
	"om/internal/ui"
)

func Version() {
	ui.PrintBanner()
	fmt.Printf("om %s\n", constants.Version)
}
