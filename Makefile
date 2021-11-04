##
# Project MusicPlayerBackend
#
# @file
# @version 1.0

ENTRYPOINT=cmd/main.go
BINARY=server.out

all: build

build:
	go build -o ${BINARY} ${ENTRYPOINT}

run: build
	chmod +x ${BINARY}
	./${BINARY}

clean:
	go clean
	rm ${BINARY}

# end
