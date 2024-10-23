FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy the Go module files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o pokemon-api ./cmd/api/main.go

# Use a smaller base image for the final image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/pokemon-api .

# Expose the port your app will listen on (adjust if needed)
EXPOSE 8080

# Command to run your app
CMD ["./pokemon-api"]