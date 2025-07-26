# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary (static binary for alpine compatibility)
RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener ./main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener .

# Copy .env file for database, Redis, and API configuration
COPY .env .

# Expose the API port (default 8080, configurable via .env)
EXPOSE 8080

# Run the binary
CMD ["./url-shortener"]