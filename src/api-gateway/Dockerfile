# Multi-stage docker file to build Go image

# 1. Start from golang base image
FROM golang:alpine as builder

# 2. Add Maintainer info
LABEL maintainer="Iheanacho Victor nonso"

# 3. Set the current working directory inside the container
WORKDIR /app

# 4. Copy go mod and sum files into the working directory
COPY go.mod go.sum ./

# 5. Download dependencies.
RUN go mod download

# 6. Copy the source to the working directory in the container
COPY . .

# 7. Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin .

# 8. Start a new stage
FROM alpine:latest

# 11. Set working directory
WORKDIR /root/

# 12. Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin .

# 13. Command to run the executable
CMD ["./bin"]
