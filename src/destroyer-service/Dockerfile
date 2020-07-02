# Multi-stage docker file to build Go image

# 1. Start from golang base image
FROM golang:alpine as builder

# 2. Add Maintainer info
LABEL maintainer="Iheanacho Victor nonso"

# 3. Set the current working directory inside the container
WORKDIR /app

# 4. Copy go mod and sum files
COPY src/destroyer-service/go.mod src/destroyer-service/go.sum ./

# 5. Download dependencies.
RUN go mod download

# 6. Copy proto file into the working directory
COPY proto/destroyer-service.proto .

# 7. Install depenedecies for generating proto
#RUN  generate proto-file into proper directory

# 8. Generate go package from the proto file
#RUN  generate proto-file into proper directory

# 9. Copy the source to the working directory in the container
COPY src/destroyer-service/. .

# 10. Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin .

# 11. Start a new stage
FROM alpine:latest

# 12. Set working directory
WORKDIR /root/

# 13. Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin .

# 14. Command to run the executable
CMD ["./bin"]