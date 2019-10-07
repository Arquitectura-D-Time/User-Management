# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.13.1 as dev

# Install dependencies
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/go-chi/chi

# copy the local package files to the container workspace
ADD . /go/src/User-Management

# Setting up working directory
WORKDIR /go/src/User-Management

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main.

# alpine image
FROM alpine:3.9.2 as prod
# Setting up working directory
WORKDIR /root/
# copy movies binary
COPY --from=dev /go/src/User-Management .
# Service listens on port 5005.
EXPOSE 5005
# Run the movies microservice when the container starts.
ENTRYPOINT ["./usermanagement"]
