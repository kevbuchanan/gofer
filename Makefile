.PHONY: all
default: install

lib:
	go install

cmd: lib
	go install github.com/kevinbuch/gofer/cmd/gofer

install: cmd

test: lib
	@cd test && go test -v
