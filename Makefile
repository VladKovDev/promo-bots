.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/app/main.go

run: build
	./.bin/bot