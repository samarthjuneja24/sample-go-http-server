## Steps to run

#### 1. Build your Docker image.
<code>docker build -t my-go-server . </code>

#### 2. Run your Docker container, mapping port 8000 of the container to port 8000 on the host.

<code>docker run -p 8000:8000 my-go-server</code>