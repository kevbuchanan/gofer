package gofer

import (
	"errors"
	"io"
	"net/http"
	"os"
)

type Download struct {
	Target *os.File
	Source *io.ReadCloser
	Length int64
	Read   int64
	Reads  chan int64
	Errors chan error
}

func maxRead(leftToRead, defaultRead int64) int64 {
	if leftToRead < defaultRead {
		return leftToRead
	} else {
		return defaultRead
	}
}

func (download *Download) Start() {
	for {
		toRead := maxRead(download.Length-download.Read, 1000)
		nRead, err := io.CopyN(download.Target, *download.Source, toRead)
		if err != nil {
			download.Errors <- err
			break
		}
		download.Read += nRead
		download.Reads <- nRead
	}
}

func getSourceReader(source string) (*io.ReadCloser, int64, error) {
	resp, err := http.Get(source)
	if err != nil || resp.StatusCode != 200 {
		return nil, 0, errors.New("Could not locate file")
	}
	return &resp.Body, resp.ContentLength, nil
}

func NewDownload(source string, dest string) Download {
	reads := make(chan int64, 10)
	errors := make(chan error, 10)

	sourceReader, length, _ := getSourceReader(source)
	destFile, _ := os.Create(dest)

	download := Download{
		Target: destFile,
		Source: sourceReader,
		Length: length,
		Read:   0,
		Reads:  reads,
		Errors: errors,
	}

	return download
}
