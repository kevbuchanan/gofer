package main

import "strings"

func Setup(args []string) (string, string) {
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
