package gofer

import "strings"

type Setup struct {
	FileUrl     string
	Destination string
	Reads       chan int64
	Errors      chan error
}

func parseArgs(args []string) (string, string) {
	fileUrl := args[1]
	var dest string
	if len(args) > 2 {
		dest = args[2]
	} else {
		parts := strings.Split(fileUrl, "/")
		dest = parts[len(parts)-1]
	}
	return fileUrl, dest
}

func NewSetup(args []string) Setup {
	fileUrl, dest := parseArgs(args)

	setup := Setup{
		FileUrl:     fileUrl,
		Destination: dest,
		Reads:       make(chan int64, 10),
		Errors:      make(chan error, 10),
	}

	return setup
}
