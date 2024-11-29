FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o receipt-processor ./cmd/api/main.go

# Use a minimal image for the final container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/receipt-processor .
CMD ["./receipt-processor"]
