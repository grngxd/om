package logger

import (
	"fmt"
	"om/internal/ui/colours"

	"github.com/fatih/color"
)

func Logf(msg string, args ...any) {
	fmt.Printf(colours.Builder("[om] ", color.Faint, color.Bold)+msg, args...)
}

func Log(msg string) {
	Logf(msg + "\n")
}

func Successf(msg string, args ...any) {
	fmt.Printf(colours.Builder("[om] ", color.FgGreen, color.Bold)+"success: "+msg, args...)
}

func Success(msg string) {
	Successf(msg + "\n")
}

func Warnf(msg string, args ...any) {
	fmt.Printf(colours.Builder("[om] ", color.FgYellow, color.Bold)+"warning: "+msg, args...)
}

func Warn(msg string) {
	Warnf(msg + "\n")
}

func Errorf(msg string, args ...any) {
	fmt.Printf(colours.Builder("[om] ", color.FgRed, color.Bold)+"error: "+msg, args...)
}

func Error(msg string) {
	Errorf(msg + "\n")
}
