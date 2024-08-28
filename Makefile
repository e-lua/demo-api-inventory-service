SHELL:=/bin/bash -O extglob
BINARY=ms-api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Clean up old binaries
clean:
	rm -f ${BINARY}

#go tool commands
build:
	go build ${LDFLAGS} -o ${BINARY} main.go

