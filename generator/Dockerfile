FROM golang:1.7-alpine

ADD . /go/src/generator

RUN go install generator

ENTRYPOINT /go/bin/generator

EXPOSE 8080

