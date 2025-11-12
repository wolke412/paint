package paint

import "fmt"

// Color codes
const (
	BLACK         = 30
	RED           = 31
	GREEN         = 32
	YELLOW        = 33
	BLUE          = 34
	MAGENTA       = 35
	CYAN          = 36
	WHITE         = 37

	BRIGHT_BLACK   = 90
	BRIGHT_RED     = 91
	BRIGHT_GREEN   = 92
	BRIGHT_YELLOW  = 93
	BRIGHT_BLUE    = 94
	BRIGHT_MAGENTA = 95
	BRIGHT_CYAN    = 96
	BRIGHT_WHITE   = 97

	NOCOLOR = -1
)

func FG(color int, text string) string {
	if color == NOCOLOR {
		return text
	}

	return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

func BG(color int, text string) string {
	if color == NOCOLOR {
		return text
	}
	return fmt.Sprintf("\033[%dm%s\033[0m", color + 10, text)
}

func Both(fg int, bg int, text string) string {
	return FG(fg, BG(bg, text))
	// return fmt.Sprintf("\033[3%d;4%dm%s\033[0m", fg, bg, text)
}


func Black(text string) string        { return FG(BLACK, text) }
func Red(text string) string          { return FG(RED, text) }
func Green(text string) string        { return FG(GREEN, text) }
func Yellow(text string) string       { return FG(YELLOW, text) }
func Blue(text string) string         { return FG(BLUE, text) }
func Magenta(text string) string      { return FG(MAGENTA, text) }
func Cyan(text string) string         { return FG(CYAN, text) }
func White(text string) string        { return FG(WHITE, text) }
func BrightWhite(text string) string  { return FG(BRIGHT_WHITE, text) }
func BrightRed(text string) string    { return FG(BRIGHT_RED, text) }
func BrightGreen(text string) string  { return FG(BRIGHT_GREEN, text) }
func BrightYellow(text string) string { return FG(BRIGHT_YELLOW, text) }
func BrightBlue(text string) string   { return FG(BRIGHT_BLUE, text) }
func BrightMagenta(text string) string { return FG(BRIGHT_MAGENTA, text) }
func BrightCyan(text string) string   { return FG(BRIGHT_CYAN, text) }

