# Use an official Ubuntu image as the base
FROM ubuntu:latest

# Install necessary packages including Git and wget
RUN apt-get update && apt-get install -y \
    git \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION=1.23.0
RUN wget https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

# Set the working directory for Go
WORKDIR /go/src/app

# Copy go.mod and go.sum from the go-server directory
COPY go-server/go.mod go-server/go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the remaining source code into the container
COPY go-server .

# Copy the wait-for-it.sh script into the container
COPY go-server/wait-for-it.sh /app/wait-for-it.sh

# Make the wait-for-it.sh script executable
RUN chmod +x /app/wait-for-it.sh

# Build the Go application
RUN go build -o /go/bin/app .

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the application
CMD ["/app/wait-for-it.sh", "mysql:3306", "--", "/go/bin/app"]
