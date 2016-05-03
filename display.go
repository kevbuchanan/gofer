package gofer

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
)

type Display struct {
	signals chan os.Signal
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func (display Display) Status(percentDone int) {
	fmt.Print("\r")
	fmt.Print("[")
	fmt.Print(strings.Repeat("=", percentDone/2))
	fmt.Print(">")
	fmt.Print(strings.Repeat(" ", 50-percentDone/2))
	fmt.Print("]")
	fmt.Printf("(%d%%)", percentDone)
}

func (display Display) Done() {
	showCursor()
	fmt.Println()
	os.Exit(0)
}

func (display Display) Error(e error) {
	showCursor()
	fmt.Println(e)
	os.Exit(1)
}

func handleInterrupt(signals chan os.Signal) {
	for _ = range signals {
		fmt.Println()
		showCursor()
		os.Exit(1)
	}
}

func (display Display) handleCursor() {
	hideCursor()
	signal.Notify(display.signals, os.Interrupt)
	signal.Notify(display.signals, os.Kill)
	go handleInterrupt(display.signals)
}

func NewDisplay() Display {
	display := Display{
		signals: make(chan os.Signal, 1),
	}
	display.handleCursor()
	return display
}
