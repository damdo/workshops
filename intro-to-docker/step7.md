But were did we left our app backend?
Well now that we have got some experience with starting, inspecting and stopping containers we can spin up an instance of our app backend.
We know the drill, it is a run, this though, we want to also do some port magic.
By default our app backend is listening on port 8081. Athough in order to be able to reach the app on that conainer port we need to bind our host machine port to the one of the container we are interested in.
In this case we are going to map our host port 8082 to the port 8081 of the container.
We can do this by specifying the `-p` flag like so:
`docker run -d --name my-backend -p 8082:8081 backend`

We can verify that the container started correctly by issuing a ps:
`docker ps`{{execute}}

On this ps we can also clearly see how the port mapping is set up.
`8082->8081`

We can log what is happening on our container to confirm that everything has been started correctly
`docker logs my-backend`{{execute}}
