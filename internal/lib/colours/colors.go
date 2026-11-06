package colours

import "github.com/fatih/color"

func Green(s string) string {
	return color.New(color.FgGreen).SprintFunc()(s)
}

func Yellow(s string) string {
	return color.New(color.FgYellow).SprintFunc()(s)
}

func Red(s string) string {
	return color.New(color.FgRed).SprintFunc()(s)
}

func Faint(s string) string {
	return color.New(color.Faint).SprintFunc()(s)
}
