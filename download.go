package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Download struct {
	Target *os.File
	Source *io.ReadCloser
	Length int64
	Done   bool
	Error  bool
}

type Progress struct {
	Download *Download
	Percent  int
	Error    bool
}

func (progress *Progress) Update() {
	for progress.Percent < 100 {
		info, err := progress.Download.Target.Stat()
		if err != nil {
			progress.Download.Error = true
		}
		soFar := info.Size()
		progress.Percent = int(soFar * 100 / progress.Download.Length)
		time.Sleep(100 * time.Millisecond)
	}
}

func (progress *Progress) Watch() {
	for !progress.Download.Done {
		if progress.Download.Error {
			fmt.Println("Error")
			os.Exit(1)
		}
		display(progress.Percent)
		time.Sleep(100 * time.Millisecond)
	}
	display(100)
	fmt.Println()
}

func (download *Download) Start() {
	_, err := io.Copy(download.Target, *download.Source)
	if err != nil {
		download.Error = true
	}
	download.Done = true
}

func display(percentDone int) {
	fmt.Print("\r")
	fmt.Print("[")
	fmt.Print(strings.Repeat("=", percentDone/2))
	fmt.Print(">")
	fmt.Print(strings.Repeat(" ", 50-percentDone/2))
	fmt.Print("]")
	fmt.Printf("(%d%%)", percentDone)
}

func requestFile(source string) *http.Response {
	resp, err := http.Get(source)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error:", "File not found")
		os.Exit(1)
	}
	return resp
}

func setup(args []string) (string, string) {
	file := os.Args[1]
	var dest string
	if len(os.Args) > 2 {
		dest = os.Args[2]
	} else {
		parts := strings.Split(file, "/")
		dest = parts[len(parts)-1]
	}
	return file, dest
}

func main() {
	file, dest := setup(os.Args)
	resp := requestFile(file)
	destFile, _ := os.Create(dest)

	download := Download{
		Source: &resp.Body,
		Target: destFile,
		Length: resp.ContentLength,
		Done:   false,
	}

	progress := Progress{
		Download: &download,
		Percent:  0,
	}

	go progress.Update()
	go download.Start()
	progress.Watch()
}
