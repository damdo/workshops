Container images are composable (stackable) blocks of layers.
![layers](https://i.stack.imgur.com/fotPN.jpg)

credits: https://stackoverflow.com/questions/55174274/understanding-docker-layers-and-future-changes

![layers2](https://miro.medium.com/max/1400/1*hZgRPWerZVbaGT8jJiJZVQ.jpeg)

credits: https://medium.com/docker-captain/docker-basics-f1a06fde18fb

A Dockerfile declares how a container image is supposed to be formed and built.
It vaguely recalls how we setup software on a machine:
we specify the steps to copy files around,
we install a package or maybe even build it from source.

Let's suppose we have an app with a backend and a frontend. The backend is in the `backend/` folder and the frontend is in the `frontend/` folder. We can see them [here](https://github.com/damdo/workshops/tree/master/intro-to-docker/assets/).

Let's now suppose we want to build a container image for the backend.

To do that we can define a Dockerfile in the `backend/` dir that looks like this:

<pre class="file" data-filename="Dockerfile" data-target="replace">

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
</pre>

Now, paste above content—see the copy icon in the upper-right corner of above block—and it will be automatically saved in a file called `Dockerfile` in the `backend/` directory.
