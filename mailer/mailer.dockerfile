# Base Go Image
FROM golang:1.23.4-alpine AS builder

# Set working directory
WORKDIR /app

# Add source code
COPY . /app

# Build the binary and add environment variable through CGO_ENABLED
RUN CGO_ENABLED=0 go build -o mailer ./cmd/event

RUN chmod +x /app/mailer

# Build a small image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder app/mailer ./

# Command to run the executable
CMD ["./mailer"]
