# syntax=docker/dockerfile:1

# tell Docker what base image you would like to use for your application
FROM golang:1.19

# Set destination for COPY
# To make things easier when running the rest of your commands, create a directory inside the image that you're building. 
# This also instructs Docker to use this directory as the default destination for all subsequent commands.
# This way you don't have to type out full file paths in the Dockerfile, the relative paths will be based on this directory.
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD [ "/docker-gs-ping" ]
