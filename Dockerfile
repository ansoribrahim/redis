# Use the official Go image as the base
FROM golang:1.17-alpine3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

RUN go build -o main .

# Set the entrypoint to the built Go application
CMD ["./main"]