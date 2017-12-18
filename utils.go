package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// Position : Position coordinates
type Position struct {
	line   int
	column int
}

// Soma : Change position by summing displacement and position
func (pos1 *Position) Soma(pos2 Position) Position {
	return Position{pos1.line + pos2.line, pos1.line + pos2.line}
}

// Init : Set up terminal for rendering the game
func Init() {
	rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()
}

// Quit : Finish and restore terminal
func Quit() {
	rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()
}

// ESC : Escape key
const ESC = "\x1b"

// Refresh : Clear screen and move cursor back to origin
func Refresh() {
	fmt.Printf("%s[2J", ESC)
	MoveCursor(Position{0, 0})
}

// MoveCursor : Move cursor to `p` position coordinates
func MoveCursor(p Position) {
	fmt.Printf("%s[%d;%df", ESC, p.line+1, p.line+1)
}

// FgRed : Return red foreground color for terminal
func FgRed(s string) string { return ansi(31, s) }

// FgGreen : Return green foreground color for terminal
func fgGreen(s string) string { return ansi(32, s) }

// FgBlue : Return blue foreground color for terminal
func FgBlue(s string) string { return ansi(34, s) }

// BgRed : Return red background color for terminal
func BgRed(s string) string { return ansi(41, s) }

// BgGreen : Return green background color for terminal
func BgGreen(s string) string { return ansi(42, s) }

// BgBlue : Return blue background color for terminal
func BgBlue(s string) string { return ansi(44, s) }

// Bright : Return bright colors for terminal
func Bright(s string) string { return ansi(1, s) }

var ansiRE *regexp.Regexp

func ansi(code int, s string) string {
	if ansiRE == nil {
		ansiRE = regexp.MustCompile(`^` + ESC + `\[(\d+(?:;\d+)*m.*` + ESC + `\[0m)$`)
	}

	parts := ansiRE.FindStringSubmatch(s)
	if parts == nil {
		return fmt.Sprintf("%s[%dm%s%s[0m", ESC, code, s, ESC)
	} else {
		return fmt.Sprintf("%s[%d;%s", ESC, code, parts[1])
	}

	return ""
}
