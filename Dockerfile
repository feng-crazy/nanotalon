# Use the official Golang image to create a build artifact
FROM golang:1.24-alpine AS builder

# Install git (needed for go mod download)
RUN apk add --no-cache git ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o nanotalon ./cmd/main.go

# Use a minimal base image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -g 65532 nonroot &&\
    adduser -D -u 65532 -G nonroot nonroot

# Create workspace directory
RUN mkdir -p /workspace && chown nonroot:nonroot /workspace

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/nanotalon /nanotalon

# Make the binary executable
RUN chmod +x /nanotalon

# Expose port 18790 (default for nanotalon)
EXPOSE 18790

# Change ownership of the binary to nonroot user
RUN chown nonroot:nonroot /nanotalon

# Switch to non-root user
USER nonroot:nonroot

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD /nanotalon status || exit 1

# Command to run the executable
ENTRYPOINT ["/nanotalon"]
CMD ["--help"]