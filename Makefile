##
# Project MusicPlayerBackend
#
# @file
# @version 1.0

ENRTYPOINT=cmd/main.go
BINARY=server.out

all: build

build:
	go build ${ENTRYPOINT} -o ${BINARY}

run: build
	./${BINARY}

clean:
	go clean
	rm ${BINARY}

# end
