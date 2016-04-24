package test

import (
	"github.com/kevinbuch/gofer"
	"testing"
)

func TestSetupUrl(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url}
	fileUrl, _ := gofer.Setup(args)

	if fileUrl != url {
		t.Errorf("Expected %s but was %s", url, fileUrl)
	}
}

func TestSetupDefaultDest(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url}
	_, dest := gofer.Setup(args)

	if dest != "1.txt" {
		t.Errorf("Expected %s but was %s", "1.txt", dest)
	}
}

func TestSetupDest(t *testing.T) {
	url := "http://files.com/1.txt"
	args := []string{"gofer", url, "file.txt"}
	_, dest := gofer.Setup(args)

	if dest != "file.txt" {
		t.Errorf("Expected %s but was %s", dest, "file.txt")
	}
}
