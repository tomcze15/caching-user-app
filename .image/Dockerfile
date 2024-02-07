FROM golang:1.21

RUN go install github.com/cosmtrek/air@latest

WORKDIR /usr/src/app

COPY ./app/go.mod .

RUN go mod tidy

COPY ./app .