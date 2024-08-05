FROM docker.io/golang:1.21-alpine

WORKDIR /app

COPY ./go-app /app/go-app
RUN chmod +x /app/go-app

ENTRYPOINT ["/app/go-app"]
