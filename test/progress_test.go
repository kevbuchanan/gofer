package test

import (
	"github.com/kevinbuch/gofer"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestProgress(t *testing.T) {
	prop := func(fileSize int) bool {
		display := gofer.NewDisplay()
		statuses := make(chan int, 101)
		reads := make(chan int64, 101)
		size := int64(fileSize / 100)
		for i := 1; i < 100; i++ {
			reads <- size
		}
		reads <- -1
		progress := gofer.Progress{
			Length:  int64(fileSize),
			Reads:   reads,
			Errors:  make(<-chan error),
			Status:  statuses,
			Display: display,
		}

		progress.Update()

		success := true
		for percent := range statuses {
			if percent == -1 {
				break
			}
			if percent < 0 || percent > 100 {
				success = false
			}
		}

		display.Reset()

		return success
	}

	positiveInt := func(vals []reflect.Value, r *rand.Rand) {
		v := reflect.New(reflect.TypeOf(1)).Elem()
		posint := rand.Int63()
		v.SetInt(posint)
		vals[0] = v
	}

	config := quick.Config{
		Values: positiveInt,
	}

	if err := quick.Check(prop, &config); err != nil {
		t.Error(err)
	}
}
