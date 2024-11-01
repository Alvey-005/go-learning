# syntax=docker/dockerfile:1
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Air for hot-reloading
RUN go install github.com/cosmtrek/air@v1.40.4

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the container
COPY . .

# Copy the air config file
COPY .air.toml ./

# Expose the application port
EXPOSE 5000

# Run the application with Air
CMD ["air"]
