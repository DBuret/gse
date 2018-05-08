# GSE

A small standalone web app that just display its env &amp; information about the HTTP request it received

Its goal is to be embedded in a minimal container, to help us debugging the setup of our containers orchestrator & reverse proxies.

**This mini project was mainly a first contact with golang and an attempt to build container image from scratch, this is not a real project...** but it works ;-)

## Compile
Goal is to create a small but standalone binary to allow us to build a small container image. 

So we ask for static linking 

```
$ GGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-w" .
$ strip gse
```

The resulting binary is 4 Mb 

Run GSE:
```
$ ./gse
```
point your web browser to http://localhost:28657/gse => It works.

## Create docker image
The idea is to create an image from scratch.

Since the app has been statically linked, the Dockerfile is simply
```
FROM scratch
ADD gse /
CMD ["/gse"]
```

Build image (result: 4Mb)
```
$ sudo docker build -t gse .
```


Run locally with docker
```
$ sudo docker run -p 28657:28657 gse
```

point your web browser to http://localhost:28657/gse => working.

## Use gse
GSE will by default answer on URL path = /gse and port 28657.
You can change this behaviour by using the env vars GSE_BASEPATH and GSE_PORT

