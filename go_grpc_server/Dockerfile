# Use the official Golang image as the base image
FROM golang:1.21.3

# Set the working directory in the container
WORKDIR /app

# Copy the Go application code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o mygrpcserver

# Expose the port that your gRPC server listens on
EXPOSE 8080

# Define the command to run your gRPC server when the container starts
CMD ["./mygrpcserver"]

