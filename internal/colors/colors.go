package colors

import "fmt"

var (
	ansiColors = map[string]string{
		"black":   "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	}

	lookup = map[string]Colorer{
		"black":   Black,
		"red":     Red,
		"green":   Green,
		"yellow":  Yellow,
		"blue":    Blue,
		"magenta": Magenta,
		"cyan":    Cyan,
		"white":   White,
	}
)

type Colorer func(string) string

func Get(s string) Colorer {
	return lookup[s]
}

func Black(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["black"], ansiColors["black"]), s)
}

func Red(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["red"], ansiColors["red"]), s)
}

func Green(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["green"], ansiColors["green"]), s)
}

func Yellow(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["yellow"], ansiColors["yellow"]), s)
}

func Blue(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["blue"], ansiColors["blue"]), s)
}

func Magenta(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["magenta"], ansiColors["magenta"]), s)
}

func Cyan(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["cyan"], ansiColors["cyan"]), s)
}

func White(s string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[%sm%%s\033[%sm", ansiColors["white"], ansiColors["white"]), s)
}
