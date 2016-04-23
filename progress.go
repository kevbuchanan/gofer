package main

import (
	"fmt"
	"os"
	"time"
)

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
			os.Exit(1)
		}
		Display(progress.Percent)
		time.Sleep(100 * time.Millisecond)
	}
	Display(100)
	fmt.Println()
}

func NewProgress(download *Download) Progress {
	progress := Progress{
		Download: download,
		Percent:  0,
	}

	return progress
}
