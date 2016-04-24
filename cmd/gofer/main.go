package main

import (
	"github.com/kevinbuch/gofer"
	"os"
)

func main() {
	fileUrl, dest := gofer.Setup(os.Args)

	download := gofer.NewDownload(fileUrl, dest)
	progress := gofer.NewProgress(&download)

	go progress.Update()
	go download.Start()
	progress.Watch()
}
