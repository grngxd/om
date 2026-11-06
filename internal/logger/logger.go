package logger

import (
	"fmt"
)

func Logf(msg string, args ...any) {
	fmt.Printf("[om] "+msg, args...)
}

func Log(msg string) {
	Logf(msg + "\n")
}

func Errorf(msg string, args ...any) {
	fmt.Printf("[om] error: "+msg, args...)
}

func Error(msg string) {
	Errorf(msg + "\n")
}
