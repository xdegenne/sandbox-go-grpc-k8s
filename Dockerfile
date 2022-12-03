# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build
WORKDIR /app

RUN apk update && apk add git

COPY go.mod go.sum main.go ./
COPY pkg pkg

RUN go mod download
RUN go build

FROM alpine:3.15

WORKDIR /app
RUN adduser -u 999 -S -h /app kosmos
RUN mkdir /var/run/plugin && chown 999 /var/run/plugin
RUN apk update

USER 999

COPY --from=build /app/hello /app/hello

CMD [ "/app/hello" ]
