package logger

import (
	"fmt"
	"om/internal/lib/colours"
)

func Logf(msg string, args ...any) {
	fmt.Printf(colours.Faint("[om] ")+msg, args...)
}

func Log(msg string) {
	Logf(msg + "\n")
}

func Successf(msg string, args ...any) {
	fmt.Printf(colours.Green("[om] ")+"success: "+msg, args...)
}

func Success(msg string) {
	Successf(msg + "\n")
}

func Warnf(msg string, args ...any) {
	fmt.Printf(colours.Yellow("[om] ")+"warning: "+msg, args...)
}

func Warn(msg string) {
	Warnf(msg + "\n")
}

func Errorf(msg string, args ...any) {
	fmt.Printf(colours.Red("[om] ")+"error: "+msg, args...)
}

func Error(msg string) {
	Errorf(msg + "\n")
}
