package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Download struct {
	Target *os.File
	Source *io.ReadCloser
	Length int64
	Done   bool
	Error  bool
}

func (download *Download) Start() {
	_, err := io.Copy(download.Target, *download.Source)
	if err != nil {
		download.Error = true
	}
	download.Done = true
}

func NewDownload(source string, dest string) Download {
	resp, err := http.Get(source)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error:", "File not found")
		os.Exit(1)
	}

	destFile, _ := os.Create(dest)

	download := Download{
		Source: &resp.Body,
		Target: destFile,
		Length: resp.ContentLength,
		Done:   false,
	}
	return download
}
