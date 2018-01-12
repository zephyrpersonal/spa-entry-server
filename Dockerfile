FROM golang:alpine

RUN pwd

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go get -t github.com/valyala/fasthttp

ADD server.go src

ENV GOBIN $GOPATH/bin

RUN go install src/server.go