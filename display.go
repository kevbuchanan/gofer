package main

import (
	"fmt"
	"strings"
)

func Display(percentDone int) {
	fmt.Print("\r")
	fmt.Print("[")
	fmt.Print(strings.Repeat("=", percentDone/2))
	fmt.Print(">")
	fmt.Print(strings.Repeat(" ", 50-percentDone/2))
	fmt.Print("]")
	fmt.Printf("(%d%%)", percentDone)
}
