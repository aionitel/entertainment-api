# Use the official Golang image as the base image
FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=3004

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code to the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the application will listen on
EXPOSE $PORT

# Run app
CMD ["./main"]