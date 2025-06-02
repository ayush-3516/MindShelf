FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o mindshelf-api ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -g '' appuser

# Copy the binary from builder stage and set correct ownership
COPY --from=builder /app/mindshelf-api /app/mindshelf-api
RUN chown appuser:appuser /app/mindshelf-api

USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Command to run
CMD ["/app/mindshelf-api"]