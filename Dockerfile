# Use a minimal base image for Go apps
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build .

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app listens on
EXPOSE 8080

# Run the binary
CMD ["./htmx-fr"]
