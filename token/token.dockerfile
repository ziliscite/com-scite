# Base Go Image
FROM golang:1.23.4-alpine AS builder

# Set working directory
WORKDIR /app

# Add source code
COPY . /app

# Build the binary and add environment variable through CGO_ENABLED
RUN CGO_ENABLED=0 go build -o token ./cmd/rpc

RUN chmod +x /app/token

# Build a small image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder app/token ./

# Copy migrations files
COPY migrations ./migrations

# Expose gRPC port
EXPOSE 50051

# Command to run the executable
CMD ["./token"]
