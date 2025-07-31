FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]