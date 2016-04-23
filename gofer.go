package main

import "os"

func main() {
	fileUrl, dest := Setup(os.Args)

	download := NewDownload(fileUrl, dest)
	progress := NewProgress(&download)

	go progress.Update()
	go download.Start()
	progress.Watch()
}
