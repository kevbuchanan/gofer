package gofer

import (
	"fmt"
	"os"
	"strings"
)

func Display(percentDone int) {
	fmt.Print("\033[?25l") // Hide cursor
	fmt.Print("\r")
	fmt.Print("[")
	fmt.Print(strings.Repeat("=", percentDone/2))
	fmt.Print(">")
	fmt.Print(strings.Repeat(" ", 50-percentDone/2))
	fmt.Print("]")
	fmt.Printf("(%d%%)", percentDone)
}

func DisplayDone() {
	fmt.Print("\033[?25h") // Show cursor
	fmt.Println()
	os.Exit(0)
}

func DisplayError(e error) {
	fmt.Println(e)
	os.Exit(1)
}
