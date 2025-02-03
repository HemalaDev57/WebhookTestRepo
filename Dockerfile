# Use Golang image as base
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests to the container
COPY app/go.mod app/go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY app/ .

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
