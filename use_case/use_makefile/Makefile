export tag=v1.0
root:
	export ROOT=go101/use_case/use_makefile

build:
	echo "building httpserver"
	makedir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "releasing httpserver"

push: release
	echo "pushing httpserver"