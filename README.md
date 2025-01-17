# ContainerHowTo
a simple example walk thru of how to create a service and use containers on podman and kubernetes

## Description of Chapters
this project is broken down into chapters of increasing complexity.  To start this howto, switch to the
first chapter branch in the repo with:
```
git switch chap-1
```

## Requirements
### We assume that you already have the following setup and working
1. podman 5,  setup rootless ( we use 5.2.2 )  

1. RKE2, kubernetes cluster  ( we use v1.28.15+rke2r1 )  

1. helm v4.0  

1. access to docker.io or a local container image registry  
 For a local container image repo we run the image  
    docker.io/library/registry:2

1. make tools

1. go 1.21  (we use 1.23)  

### go knowledge is not needed
go is used to make the server application as it is simple and has no external dependices. You do not 
need to understand go to go thru this howto

