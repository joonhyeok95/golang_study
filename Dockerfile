FROM golang:1.21-alpine

WORKDIR /app

COPY ./go-app /app/main

ENTRYPOINT ["./main"]
