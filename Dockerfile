# 1. Build stage: compile the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# 2. Runtime stage: super-tiny final image
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/app .
COPY templates ./templates
COPY static ./static

EXPOSE 8080
CMD ["./app"]