# Start from a lightweight version of the Go image
FROM golang:alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
# CGO_ENABLED=0: disables CGO (C Go) to compile the Go app without C library
# GOOS=linux: compiles the Go app for Linux
# go build -o main: compiles the Go app and names the executable "main"
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# If you are using ARM architecture, you can use the following command:
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the data directory into the container
# Otherwise, the data directory will not be found
# Because the data directory is not included in the
# binary file when building the Go app
COPY --from=builder /app/data/ ./data/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
