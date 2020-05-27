Ok, great!
It is now time to create a second image for our frontend.
Our frontend is super simple this time, we have just an index.html file with some javascript on it.
The javascript is going to fetch some data from our backend and return it out on the frontend.

To serve our HTML  we can use whatever web server engine we prefer, in this case a good choice is NGINX.

We just need to specify it as a base image and put our HTML in the default folder NGINX serves, which is `/usr/share/nginx/html`.

To achive this we can define a Dockerfile in the `frontend/` dir that looks like this:

<pre class="file" data-filename="frontend/Dockerfile" data-target="replace">
FROM nginx:latest

COPY index.html /usr/share/nginx/html
</pre>

Now, paste above content—see the copy icon in the upper-right corner of above block—and it will be automatically saved in a file called `Dockerfile` in the `frontend/` directory.

Let's then build a `frontend` image off of that:
`docker build -t frontend .`{{execute}}
