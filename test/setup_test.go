package test

import (
	"github.com/kevinbuch/gofer"
	"testing"
)

func TestSetupUrl(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url}
	setup := gofer.NewSetup(args)

	if setup.FileUrl != url {
		t.Errorf("Expected %s but was %s", url, setup.FileUrl)
	}
}

func TestSetupDefaultDest(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url}
	setup := gofer.NewSetup(args)

	if setup.Destination != "1.txt" {
		t.Errorf("Expected %s but was %s", "1.txt", setup.Destination)
	}
}

func TestSetupDest(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url, "file.txt"}
	setup := gofer.NewSetup(args)

	if setup.Destination != "file.txt" {
		t.Errorf("Expected %s but was %s", "file.txt", setup.Destination)
	}
}
