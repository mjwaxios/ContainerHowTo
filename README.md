# ContainerHowTo
a simple example walk thru of how to create a service and use containers on podman and kubernetes

## Chapter Three -- Makefile
For this chapter we will add a makefile to help us out.  
The makefile contains just a few targets for now:  
1. serv - this will build the serv program
1. image - this will build the serv container image
1. clean - this will erase the serv program 


To build the service:
```
make 
```

To build the container image:  
```
make image
```


## Next Chapter
```
git switch chap-4
```
