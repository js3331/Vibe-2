FROM golang:1.21-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and source files
COPY . .

# Download dependencies and generate go.sum
RUN go mod download
RUN go mod tidy

# Build the application
RUN go build -o /app/main ./cmd/main.go

# Verify the binary exists
RUN ls -l /app/main

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
