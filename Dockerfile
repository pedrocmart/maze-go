# Build binary.
FROM golang:1.16.0-alpine3.13 AS build-env
RUN apk add --update --no-cache git openssh
RUN go get -u github.com/gobuffalo/packr/packr

COPY . /go/src/github.com/pedrocmart/maze-go
WORKDIR /go/src/github.com/pedrocmart/maze-go/cmd
RUN packr build -o dist/goapp

# Build image.
FROM alpine:3.13
RUN  apk update && apk upgrade \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates
WORKDIR /app
COPY --from=build-env /go/src/github.com/pedrocmart/maze-go/cmd/dist/goapp /app/
ENTRYPOINT ./goapp
