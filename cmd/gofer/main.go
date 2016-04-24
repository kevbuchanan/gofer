package main

import (
	"github.com/kevinbuch/gofer"
	"os"
)

func main() {
	setup := gofer.NewSetup(os.Args)

	download := gofer.NewDownload(setup)
	progress := gofer.NewProgress(setup, download)

	go progress.Update()
	go download.Start()
	progress.Watch()
}
