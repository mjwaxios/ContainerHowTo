# ContainerHowTo
a simple example walk thru of how to create a service and use containers on podman and kubernetes

## Description of Chapters
this project is broken down into chapters of increasing complexity.  To start this howto, switch to the
first chapter branch in the repo with:
```
git switch chap-1
```

## Requirements
We use:  
 go 1.21  
 podman 5  (5.2.2)  
 RKE2 v1.28.15+rke2r1  
 helm v4.0  

 For a local container image repo we run the image  
    docker.io/library/registry:2