Great work,
now that we have declared what the image should look like we now need to build it.
In order to build an image from a Dockerfile we can use the `build` command.
The syntax is quite easy:
```
Usage:  docker build [OPTIONS] PATH | URL | -`
```

So we are going to call our image: `backend` and the Dockerfile declaration with the instructions to build it is in the `backend/` path. So let's compose our command that will start a build:

`docker build --tag backend backend/`{{execute}}

Now that we built an image we can list it together with all the other images that we have available on our machine through the `images` command:

`docker images`{{execute}}

With this command we can get the size, names (tags), id and creation date of the image.

A single image can be referred through different names and tags.
For example we can re-tag our image once we have built it and give it an alternative name.

`docker tag backend alt-backend`{{execute}}

If we now list the images again we can see that a new image is available.
`docker images`{{execute}}

NB. the image is still going to be the same one (this can be seen by the fact that they have the same IMAGE ID) but the command is reporting that it can be referred through another name.

