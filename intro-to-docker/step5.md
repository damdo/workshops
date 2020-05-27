We can see all the currently running containers on our machine by typing:
`docker ps`{{execute}}

And there it is, redis is there. The ps command also tells us:
- the CONTAINER ID, a random id that uniquely identifies the container on the machine
- the COMMAND, that triggered the process currently running in the container foreground
- the CREATED label, it tells us when the container was started
- the STATUS of the container. "Up" in this case
- the PORTS to which the container is binding internally and externally
- the NAME (which is randomly generated if not specified with the `--name` flag at `run` invocation time)

We can check the container resource utilization by running:
`docker stats --no-stream my-redis`{{execute}}

And we can check the container logs by doing:
`docker logs my-redis`{{execute}}

