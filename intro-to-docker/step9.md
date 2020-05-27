So now that the frontend image is ready it is now time to spin up a frontend instance.

Also this time we need to map the container port to an host port. Node.js by default is going to listen to the container 80 port, we then need to map for example port 80 of our host to the 80 of the container.
So `80->80`.

We will also need to link our new frontend container to the previously built my-backend container so that they can reach eachother, we can do that with the `--link` directive:
`docker run -d --rm -it -p 80:80 --link my-backend --name=my-frontend frontend`{{execute}}

Ok we should be able to see a running frontend if we check via ps:
`docker ps`

If that's the case we should be now be finally able to visit our frontend-url on port 8090 and be able to see our system fully working. It is possible to do that by clicking on the + symbol, and select  `View HTTP port 80 on Host 1`.
