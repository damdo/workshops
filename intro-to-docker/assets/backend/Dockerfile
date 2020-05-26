# The FROM instruction initializes a new build stage
# and sets the Base Image for subsequent instructions.
FROM golang:1.14.3-alpine

# The WORKDIR instruction sets the working directory
# for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions
# If the WORKDIR doesn’t exist, it will be created.
WORKDIR /go/src/app

# The COPY instruction copies new files or directories
# from <src> and adds them to the filesystem
# of the container at the path <dest>.
COPY . .

# The RUN instruction will execute any commands in
# a new layer on top of the current image and commit the results.
# The resulting committed image will be used
# for the next step in the Dockerfile
RUN go build -v ./...
RUN cp backend /usr/local/bin

# The main purpose of a CMD is to provide defaults
# for an executing container.
# These defaults can include an executable
CMD ["/usr/local/bin/backend"]
