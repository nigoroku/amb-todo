FROM golang:1.14.10-alpine3.12

WORKDIR /go/src/github.com/nigoroku/amb-todo
ADD . /go/src/github.com/nigoroku/amb-todo

ENV GO111MODULE=on

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh

EXPOSE 8082

CMD ["fresh"]