ARG PLATFORM=linux/amd64

FROM --platform=$PLATFORM golang:1.22-alpine

RUN apk add --no-cache --update \
    git \
    supervisor \
    gcc \
    musl-dev \
    openssh-client \
    gettext

ENV GOCACHE=/tmp/gocache

RUN mkdir /var/run/supervisor /var/log/supervisor

#configure golang environment variables
RUN #go env -w GOPRIVATE=bucket.digitalarsenal.net
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
RUN go install github.com/vektra/mockery/v2@latest
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
