FROM golang:1.23.4 AS builder

ADD . /app/
WORKDIR /app

RUN go mod tidy -v

RUN go build -o ./build/main ./cmd/app/main.go

FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/config config
COPY --from=builder /app/migrations/ migrations
COPY --from=builder /app/build/main  ./cmd/main

EXPOSE 8080

ENTRYPOINT ["./cmd/main"]
