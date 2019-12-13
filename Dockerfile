FROM golang:latest

LABEL maintainer="Kuramshin Rustam <kuramshin.py@yandex.ru>"

WORKDIR /go/src/mortgage

ADD ./cmd ./cmd/

ENV GOPATH /go

RUN apt-get update \
    && export GOPATH=${GOPATH} \
    && go get -d -v ./... \
    && go install -v ./...

EXPOSE 8000

CMD [ "go", "run", "main.go" ]


