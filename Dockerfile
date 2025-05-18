FROM golang:1.21-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
