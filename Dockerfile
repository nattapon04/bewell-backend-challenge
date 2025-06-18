# Stage 1: Build
FROM golang:1.23.8-alpine AS builder

# Set working directory
WORKDIR /app

# Copy Go modules definition
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app from cmd/main.go
RUN go build -o main ./cmd

# Stage 2: Run
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/main .

COPY .env ./

# Run the binary
CMD ["./main"]
