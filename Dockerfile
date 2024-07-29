# Use the official Golang image to create a build artifact.
FROM golang:1.18 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the rest of your application's source code
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Copy the binary file from the builder stage
COPY --from=builder /app/main .

# Make the "main" binary executable
RUN chmod +x main

# Command to run the executable
CMD ["./main"]
