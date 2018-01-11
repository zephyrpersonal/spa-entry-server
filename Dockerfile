FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -t github.com/valyala/fasthttp

ADD server.go $GOPATH/src

WORKDIR $GOPATH/src

RUN go build server.go