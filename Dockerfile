FROM goland/alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get https://github.com/cespare/reflex 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go built -o ./run .