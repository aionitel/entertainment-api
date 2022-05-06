FROM goland/alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get https://github.com/cespare/reflex 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go built -o ./run .

FROM alpine:latest

RUN apk --no-cache add ca-certificate

WORKDIR /root/

# copy built binary from builder to docker directory
COPY --from=builder /go/src/app/run .

EXPOSE 8080

CMD ["./run"]