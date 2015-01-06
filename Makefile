PKG:=github.com/docker/docker-network
PKG_PATH:=/go/src/$(PKG)
RUN_CMD:=docker run --privileged --rm -v "$(shell pwd)":$(PKG_PATH) -w $(PKG_PATH) golang:1.4 go

build:
	$(RUN_CMD) build -v

test:
	$(RUN_CMD) test -v ./...
