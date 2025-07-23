# 1. Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

# 2. Run stage
FROM alpine:latest
WORKDIR /
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]