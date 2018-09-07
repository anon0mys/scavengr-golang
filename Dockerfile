FROM golang:1.10

RUN mkdir -p /go/src/scavengr-golang
WORKDIR /go/src/scavengr-golang

ADD . /go/src/scavengr-golang

RUN go get -v
RUN go build
