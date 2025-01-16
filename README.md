# ContainerHowTo
a simple example walk thru of how to create a service and use containers on podman and kubernetes

## Chapter Two -- Containers
For this chapter we create a container image and run it on podman.

We add a Containerfile that is used to tell podman how to build out image.  
The Containerfile has the following line
```dockerfile
FROM scratch  
COPY serv /  
ENTRYPOINT ["/serv"]  
```
The FROM says use scratch as the base for our image.  scratch is an empty image.  
The COPY command will take our server binary we compile.  (Note that the build command has changed)  
The ENTRYPOINT tells the container engine what command to run at the start.  


To build the service we need to make a static binary image,  we do this with the CGO_ENABLED=0 flag:
```
cd serv
CGO_ENABLED=0 go build 
```

To build the container image:  
We tell podman to build the containerfile in the current directory and tag it as serv with version latest
```
podman build . -t serv:latest
```

Lets run it by telling podman to run it, map the host port 10100 to the container port 10100 and use the serv image.  
```
podman run -p 10100:10100 serv
```

Now in a web browser on this same system, goto http://localhost:10100  
It should look like:
> serv Test program  
> 1

## Next Chapter
```
git switch chap-3
```
