package gofer

type Progress struct {
	Length int64
	Reads  <-chan int64
	Errors <-chan error
	Status chan int
}

func (progress Progress) Update() {
	readSoFar := int64(0)
	for nRead := range progress.Reads {
		readSoFar += nRead
		percent := int(readSoFar * 100 / progress.Length)
		progress.Status <- percent
	}
}

func (progress Progress) Watch() {
	for {
		select {
		case percent := <-progress.Status:
			Display(percent)
			if percent == 100 {
				DisplayDone()
			}
		case downloadError := <-progress.Errors:
			DisplayError(downloadError)
		}
	}
}

func NewProgress(setup Setup, download Download) Progress {
	progress := Progress{
		Length: download.Length,
		Reads:  setup.Reads,
		Errors: setup.Errors,
		Status: make(chan int, 10),
	}

	return progress
}
