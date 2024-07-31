# syntax=docker/dockerfile:1.4
FROM golang:1.20.5-alpine AS builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

COPY . .

#RUN #--mount=type=cache,target=/go/pkg/mod/cache \
#    --mount=type=cache,target=/go-build \
#    go build -o bin/cengizhan-api main.go


CMD ["air", "-c", ".air.toml"]
