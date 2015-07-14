FROM alpine:3.2
MAINTAINER Joshua Mervine <joshua@mervine.net>

# getdown config - defaults for docker
ENV GETDOWN_ADDR     0.0.0.0
ENV GETDOWN_PORT     3000
ENV GETDOWN_BASEDIR  /data
#ENV GETDOWN_INDEX    README.md
#ENV GETDOWN_STYLE    https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
RUN mkdir -p $GOROOT $GOPATH $GOBIN /tmp/getdown ${GETDOWN_BASEDIR}

COPY . /tmp/getdown
RUN set -x; \
  apk update && apk add curl git mercurial bzr go && rm -rf /var/cache/apk/* && \
  cd /tmp/getdown && \
  go get && \
  go build -v -o getdown && \
  cp getdown $GOBIN && \
  apk del curl git mercurial bzr go && \
  cp /tmp/getdown/README.md /data/README.md && \
  rm -rf /tmp/getdown

WORKDIR $GETDOWN_BASEDIR

ENTRYPOINT ["getdown"]
# vim: ft=Dockerfile
