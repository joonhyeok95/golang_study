FROM docker.io/golang:1.21-alpine3.20

WORKDIR /app

COPY ./go-app /app/main

ENTRYPOINT ["./main"]
