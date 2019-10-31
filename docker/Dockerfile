FROM golang:1.11

ENV GO111MODULE=on

ADD . /go/src/github.com/kudejwiktor/go-api-example
WORKDIR /go/src/github.com/kudejwiktor/go-api-example

RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go install main.go
RUN ["chmod", "+x", "/go/src/github.com/kudejwiktor/go-api-example/docker/entrypoint.sh"]
RUN ["chmod", "+x", "/go/src/github.com/kudejwiktor/go-api-example/docker/wait-for-it.sh"]

ENTRYPOINT docker/entrypoint.sh

EXPOSE 8083
