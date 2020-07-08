FROM golang:1.14.4 AS build

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o app

FROM alpine:3.9
RUN apk add --update tzdata
RUN apk add --update ca-certificates
RUN apk add --update bash

COPY --from=build /app/app app
RUN chmod +x /app
