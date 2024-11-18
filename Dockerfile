# Use the official Golang image as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Run tests
RUN go test -v ./...

# Build the Go application
RUN go build -o ascii-art-web .

# Expose the port that your web server listens on
EXPOSE 8080

# Command to run the web server
CMD ["./ascii-art-web"]
