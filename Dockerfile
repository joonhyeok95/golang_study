FROM golang:alpine

WORKDIR /app

COPY ./go-app /app/main

ENTRYPOINT ["./main"]
