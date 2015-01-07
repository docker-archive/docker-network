PKG:=github.com/docker/docker-network
PKG_PATH:=/go/src/$(PKG)
IMGNAME:=dockernetwork
RUN_CMD:=docker run --privileged --rm -v "$(shell pwd)":$(PKG_PATH) -w $(PKG_PATH) $(IMGNAME)

build: dockerbuild
	$(RUN_CMD) go build -v

test: dockerbuild
	$(RUN_CMD) go test -v ./...

dockerbuild:
	docker build -t $(IMGNAME) .	
