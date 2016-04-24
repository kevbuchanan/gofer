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
	Reads  chan<- int64
	Errors chan<- error
}

func maxRead(leftToRead, defaultRead int64) int64 {
	if leftToRead < defaultRead {
		return leftToRead
	} else {
		return defaultRead
	}
}

func (download Download) Start() {
	for {
		leftToRead := download.Length - download.Read
		if leftToRead == 0 {
			break
		}
		toRead := maxRead(leftToRead, 1000)
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

func NewDownload(setup Setup) Download {
	sourceReader, length, sourceError := getSourceReader(setup.FileUrl)
	if sourceError != nil {
		setup.Errors <- sourceError
	}
	destFile, fileError := os.OpenFile(
		setup.Destination,
		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
		os.FileMode(0666),
	)
	if fileError != nil {
		setup.Errors <- fileError
	}

	download := Download{
		Target: destFile,
		Source: sourceReader,
		Length: length,
		Read:   0,
		Reads:  setup.Reads,
		Errors: setup.Errors,
	}

	return download
}
