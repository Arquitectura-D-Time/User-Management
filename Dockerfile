# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/User-Management

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/go-chi/chi

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build 


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/User-Management .

# Expose port 8080 to the outside world
EXPOSE 5003

# Command to run the executable
CMD ["./User-Management"] 