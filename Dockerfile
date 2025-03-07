# Stage 1: Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies required by SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy module files and download dependencies (cached)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build your Go binary
RUN go build -o server ./cmd/server/main.go

# Stage 2: Final minimal runtime image
FROM alpine:latest

# SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server /app/server

# If you're using a local SQLite database file, ensure it's copied into the container:
# COPY notes.db /app/notes.db

# Expose port 8080
EXPOSE 8080

# Start the application
CMD ["./server"]

