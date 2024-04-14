FROM golang:1.14.0-alpine

WORKDIR /go/src/github.com/songrgg/testservice/
COPY main.go .
CMD [ "go", "run", "main.go" ]