# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Sumit Thakur <sumitthakur769@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o sellerapp Assignment-1/myapplication.go Assignment-1/sellerapp.go

# Expose port 8000 / for internal comunication 
EXPOSE 8000

# Command to run the executable
#CMD ["./sellerapp"]
CMD ["./sellerapp","--host","0.0.0.0"]
