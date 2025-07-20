FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o dist/tracker

# ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dist/tracker .

CMD ["./tracker", "server"]