Now that our app has been developed, and since we are happy with the result we can ship our container image.

In order to ship the image we can push it to our personal image registry, to which it might be possible to hook a service that triggers a deployment.
To push our image we simply invoke the push command followed by the images

`docker push backend && docker push frontend`{{execute}}

The command will fail, to make it succeed we will need to setup and account on Docker Hub (or our favourite registry) and set it up via `docker login`
This exercise is left to the reader as practice.

FINE, The end.

Thank you for having followed this beginner workshop!
