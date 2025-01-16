# ContainerHowTo
a simple example walk thru of how to create a service and use containers on podman and kubernetes

## Chapter One
For this chapter we create a simple service that when we point a web broswer to it will answer with a increasing counter.

You can look at the source for the server but you do not need to to understand the demo by editing main.go

To build:
```
cd serv
go build
```
This will make a file called serv.  Lets run it.
```
./serv
```
Now in a web browser on this same system, goto http://localhost:10100  
It should look like:
> serv Test program  
> 1

## Next Chapter
```
git switch chap-2
```
