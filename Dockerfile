# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/project_user-management_ms

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/go-chi/chi

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build 

CMD ["./project_user-management_ms"]