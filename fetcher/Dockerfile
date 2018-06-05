FROM golang:1.7-alpine

RUN apk --update add git

ADD . /go/src/fetcher
ADD run.sh /go/bin/run.sh

# install Redis package for GO
RUN go get gopkg.in/redis.v4

# install the app
RUN go install fetcher

# set variables
ENV FETCHER_LENGTH 16
ENV FETCHER_CAPS false
ENV FETCHER_FREQ 5s

# calls /go/bin/fetcher every FETCHER_FREQ seconds
ENTRYPOINT /go/bin/run.sh
