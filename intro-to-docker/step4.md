Ok now that we got a bit familiar with container images and we created one, it is time to spin up our 
**first** container! Yay!

The command to run a new container starting from an image is the `run` command.

The syntax is fairly easy but powerful:
```
Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
```

So we can define some options (like ENV variables, ports, arguments and more) we must specify the image we want to start out container off of, and then optionally a command. If a command is not specified the one defined in CMD and/or ENTRYPOINT will be used.

`docker run --name my-redis redis`{{execute}}

Ok here we go, now we have a running redis instance!
