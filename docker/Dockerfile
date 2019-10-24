FROM golang:1.11

ENV GO111MODULE=on

ADD . /go/src/github.com/kudejwiktor/go-api-example
WORKDIR /go/src/github.com/kudejwiktor/go-api-example
#RUN go get github.com/kudejwiktor/go-api-example

RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go install app/main.go

#ENTRYPOINT docker/go-app-entrypoint.sh

EXPOSE 8080
