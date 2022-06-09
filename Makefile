APP := microserv
VERSION := $(shell git describe --tags --always --dirty)

build:
	go build -ldflags "-X main.Version='${VERSION}'" -o dist/

install:
	go install

clean:
	rm -rf dist
