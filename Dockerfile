FROM golang:1.20-buster

WORKDIR /app

RUN go mod download & go mod verify

EXPOSE 8001