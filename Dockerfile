FROM golang:1.15-alpine

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN go build -o http cmd/http/main.go

ENTRYPOINT ["./http"]
