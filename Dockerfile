# Start by selecting the base image. Here, we're using the official Go image to build our application.
FROM golang:1.20-alpine as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

# Build the Go app. You can also add flags to build a statically linked binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Use a Docker multi-stage build to create a lean production image. Start with a small base image.
FROM alpine:latest

# Add CA certificates for HTTPS requests.
RUN apk --no-cache add ca-certificates

# Set the working directory to /root/.
WORKDIR /root/

# Copy the binary from the builder stage to the production image.
COPY --from=builder /app/server .

# Expose port 8000 to the outside world.
EXPOSE 8000

# Command to run the executable.
CMD ["./server"]
