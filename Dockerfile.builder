FROM golang:1.5

RUN mkdir -p /go/src/github.com/shahriarb
ADD . /go/src/github.com/shahriarb/build_sample

WORKDIR /go/src/github.com/shahriarb/build_sample
ENV GO15VENDOREXPERIMENT 1
RUN go get
RUN go build -o build_sample

