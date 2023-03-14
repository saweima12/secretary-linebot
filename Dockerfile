FROM golang:1.20-buster

COPY . /app
WORKDIR /app

RUN go build -o ./build/secretary
EXPOSE 8080
