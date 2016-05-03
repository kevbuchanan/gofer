package gofer

import (
	"errors"
	"io"
	"net/http"
	"os"
)

type Download struct {
	Target  *os.File
	Source  *io.ReadCloser
	Length  int64
	Chunked bool
	Read    int64
	Reads   chan<- int64
	Errors  chan<- error
}

func maxRead(leftToRead, defaultRead int64) int64 {
	if leftToRead < defaultRead {
		return leftToRead
	} else {
		return defaultRead
	}
}

func readChunked(download Download) {
	_, err := io.Copy(download.Target, *download.Source)
	if err != nil {
		download.Errors <- err
	}
	download.Read += download.Length
	download.Reads <- download.Length
}

func readFixed(download Download) {
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

func (download Download) Start() {
	if download.Chunked {
		readChunked(download)
	} else {
		readFixed(download)
	}
}

func getSourceReader(source string) (*io.ReadCloser, int64, bool, error) {
	resp, err := http.Get(source)
	if err != nil || resp.StatusCode != 200 {
		return nil, 0, false, errors.New("Could not locate file")
	}
	contentLength := resp.ContentLength
	if contentLength == -1 {
		return &resp.Body, 100, true, nil
	} else {
		return &resp.Body, resp.ContentLength, false, nil
	}
}

func createFile(setup Setup) (*os.File, error) {
	return os.OpenFile(
		setup.Destination,
		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
		os.FileMode(0666),
	)
}

func NewDownload(setup Setup) Download {
	sourceReader, length, chunked, sourceError := getSourceReader(setup.FileUrl)
	if sourceError != nil {
		setup.Errors <- sourceError
	}
	destFile, fileError := createFile(setup)
	if fileError != nil {
		setup.Errors <- fileError
	}

	download := Download{
		Target:  destFile,
		Source:  sourceReader,
		Length:  length,
		Chunked: chunked,
		Read:    0,
		Reads:   setup.Reads,
		Errors:  setup.Errors,
	}

	return download
}
