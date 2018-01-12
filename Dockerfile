# build go app
FROM golang:1.9 as builder
RUN go get -t github.com/valyala/fasthttp
ADD server.go .
RUN go build server.go

#
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/server .
CMD ["./server"]