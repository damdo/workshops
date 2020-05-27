Ok, great!
It is now time to create a second image for our frontend.
Our frontend is quite simple, we have just an index.js file with some javascript on it.
The javascript is going to fetch some data from our backend and return it out on the frontend.

We just need to specify Node.js as a base image and copy our app code into the image root.

To achive this we can define a Dockerfile in the `frontend/` dir that looks like this:

<pre class="file" data-filename="frontend/Dockerfile" data-target="replace">
FROM node:alpine

COPY . .

CMD ["node", "index.js"]
</pre>

Now, paste above content—see the copy icon in the upper-right corner of above block—and it will be automatically saved in a file called `Dockerfile` in the `frontend/` directory.

Let's then build a `frontend` image off of that:
`docker build --tag my-frontend frontend/`{{execute}}
