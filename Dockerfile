# Build stage
FROM golang:1.23-alpine AS builder

# Install build tools
RUN apk add --no-cache build-base gcc git

# Set working directory
WORKDIR /app

COPY go.mod ./
RUN go mod download

# Copies everthing in folder
COPY . .

# Build and test
RUN go test -v ./...
RUN go build -o ascii-art-web . && chmod +x ascii-art-web

# Runtime stage
FROM alpine:latest

# Metadata
LABEL project="Ascii-Art-Web Dockerize"
LABEL maintainer="Toft Diederichs, Fatemeh Kheirkhah, Mahdi Kheirkhah"
LABEL description="A Dockerized ASCII Art Web Application"

# Install runtime dependencies
RUN apk add --no-cache ca-certificates bash

# Set working directory
WORKDIR /app

# Copy built files and assets
COPY --from=builder /app/ascii-art-web /app/
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
COPY --from=builder /app/banners /app/banners
# COPY --from=builder /app/. /app/.

# Expose application port
EXPOSE 8080

# Command to start the app
CMD ["./ascii-art-web"]
