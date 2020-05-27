Earlier we saw that we can use a base image to start building our image off of.

The base image is not magically appearing on our machines, but is rather pulled from a container images registry. [Docker Hub](hub.docker.com) in this case, but there are others (you can even host your own one).

On these registries we can find images that we can build on top of, but most of them are usable as is, if thatt's what just exactly what we need.

For example there are images like `redis` or `mysql` for example that are often pulled and used as is.

To find search for particular images we can use the search feature of docker:
`docker search redis`{{execute}}

and once we found one that we like we can pull it down to our machine:
`docker pull redis`{{execute}}

Pretty neat huh?
