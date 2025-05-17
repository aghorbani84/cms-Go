FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /root/

# Install necessary packages
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/app .

# Copy static files
COPY --from=builder /app/static ./static

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./app"]