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

func Bold(s string) string {
	return color.New(color.Bold).SprintFunc()(s)
}

func Builder(s string, colours ...color.Attribute) string {
	for _, c := range colours {
		s = color.New(c).SprintFunc()(s)
	}

	return color.New(colours[len(colours)-1]).SprintFunc()(s)
}
