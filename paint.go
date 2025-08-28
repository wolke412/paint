package paint

import "fmt"

// Color codes
const (
	BLACK   = 0
	RED     = 1
	GREEN   = 2
	YELLOW  = 3
	BLUE    = 4
	MAGENTA = 5
	CYAN    = 6
	WHITE   = 7
	NOCOLOR = -1
)

func Black(text string) string   { return FG(BLACK, text) }
func Red(text string) string     { return FG(RED, text) }
func Green(text string) string   { return FG(GREEN, text) }
func Yellow(text string) string  { return FG(YELLOW, text) }
func Blue(text string) string    { return FG(BLUE, text) }
func Magenta(text string) string { return FG(MAGENTA, text) }
func Cyan(text string) string    { return FG(CYAN, text) }
func White(text string) string   { return FG(WHITE, text) }

func FG(color int, text string) string {
	if color == NOCOLOR {
		return text
	}

	return fmt.Sprintf("\033[3%dm%s\033[0m", color, text)
}

func BG(color int, text string) string {
	if color == NOCOLOR {
		return text
	}
	return fmt.Sprintf("\033[4%dm%s\033[0m", color, text)
}

func Both(fg int, bg int, text string) string {
	return FG(fg, BG(bg, text))
	// return fmt.Sprintf("\033[3%d;4%dm%s\033[0m", fg, bg, text)
}
