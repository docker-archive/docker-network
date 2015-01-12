FROM golang:1.4

RUN apt-get update && apt-get install build-essential pkg-config cmake -y

ENV PKG /go/src/github.com/docker/docker-network

# Get libgit2
ENV LIBGIT2=github.com/tiborvass/git2go
ENV LIBGIT2_ORIG=github.com/libgit2/git2go
ENV PACKAGES github.com/codegangsta/cli \
  golang.org/x/sys/unix

RUN go get -d ${LIBGIT2} && \
  mkdir -p /go/src/$(dirname ${LIBGIT2_ORIG}) && \
  mv /go/src/${LIBGIT2} /go/src/${LIBGIT2_ORIG} && \
  cd /go/src/${LIBGIT2_ORIG} && \
  git checkout origin/go_backends && \
  git submodule update --init && \
  make install

RUN for i in ${PACKAGES}; do go get -v -d "$i"; done

VOLUME ${PKG}
WORKDIR ${PKG}
