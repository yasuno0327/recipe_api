FROM golang:1.11
ENV GO111MODULE on

RUN apt-get -y -qq update && \
  apt-get -y -qq install mysql-client

RUN mkdir -p /go/src/api
WORKDIR /go/src/api
ADD . /go/src/api

RUN go mod download # cache