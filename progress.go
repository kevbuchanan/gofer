package gofer

type Progress struct {
	Length  int64
	Read    int64
	Percent int
	Reads   <-chan int64
	Errors  <-chan error
	Status  chan int
}

func (progress *Progress) Update() {
	for nRead := range progress.Reads {
		progress.Read += nRead
		percent := int(progress.Read * 100 / progress.Length)
		progress.Status <- percent
	}
}

func (progress *Progress) Watch() {
	for percent := range progress.Status {
		Display(percent)
		if percent == 100 {
			DisplayDone()
			break
		}
	}
}

func NewProgress(download *Download) Progress {
	progress := Progress{
		Length: download.Length,
		Read:   0,
		Reads:  download.Reads,
		Errors: download.Errors,
		Status: make(chan int, 10),
	}

	return progress
}
