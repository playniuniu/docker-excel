# Binary name
BINARY=docker-excel
VERSION=0.1
LDFLAGS='-w -s'

build:
	# build for local os
	rm -f ./${BINARY}
	go clean
	GOPATH=${GOPATH}:`pwd` go build -o ${BINARY} -ldflags ${LDFLAGS} src/*.go

linux:
	# Build for linux
	rm -f ./${BINARY}
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} src/*.go
	upx ./${BINARY}

docker: linux
	# Build for docker
	docker build -t playniuniu/docker-excel .


# Cleans our projects: deletes binaries
clean:
	go clean
	rm -f ${BINARY}
	rm -f ${BINARY}.exe

.PHONY: build linux docker clean