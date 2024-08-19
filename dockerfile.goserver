FROM golang:1.23
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Install the dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]