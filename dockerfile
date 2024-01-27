FROM golang:1.21.3-alpine

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN apk --no-cache add curl && curl -sSf https://atlasgo.sh | sh

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download

COPY . .
