# Stage 1: Build the application with the official Go base image
FROM dockercache.mountaininfra.net/golang:1.22 as builder

# Install gettext dependency to be able to replace placesholders with real ENV variables
RUN apt-get update && \
    apt-get -y install gettext

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Copy netrc (initialization information used by the auto-login process to fetch dependencies from private repository)
COPY .netrc /root/.netrc

# Download the dependencies
RUN go mod download

# Copy the rest of application's source code
COPY cmd cmd
COPY internal internal
COPY pkg pkg

# Build Go application
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -v -o app cmd/api/*.go

# Stage 2: Create a lightweight image for the runtime
FROM alpine:latest

RUN apk add --no-cache --update \
    gettext

WORKDIR /app

COPY --from=builder /app/app .

RUN mkdir config

COPY config.dist config.dist
COPY docker/ecs/web/production/init.sh .