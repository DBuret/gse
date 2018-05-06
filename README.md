# GSE

A small standalone web app that just display its env &amp; information about the HTTP request it received

Its goal is to be embedded in a minimal container, to help us debugging the setup of our containers orchestrator & reverse proxies.

**This mini project was mainly a first contact with golang and an attempt to build container image from scratch, this is not a real project...** but it works ;-)

## Compile
Goal is to create a small but standalone binary to allow us to build a samll container image. 

So we ask for static linking (even if few libs will be missing, see docker paragraph below)

Then we strip symbolsto save a little less than 1 Mb.

```
$ GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gse gse.go
$ strip gse
```

The resulting binary is < 4mb 

Run GSE:
```
$ ./gse
```
point your web browser to http://localhost:28657/gse => It works.

## Create docker image
The idea is to create an image from scratch.

ldd against our binary shows that 3 shared libs are needed.
* glibc is not supposed to be statically linked
* ld is needed to link glibc
* regarding libpthread, I have to dig into Go threads management to check if we can avoid it (it's possible with rust)

So for the timebeing, let's just add these 3 libs to our container image, hence the Dockerfile:
```
FROM scratch
ADD libpthread.so.0 /lib64/libpthread.so.0
ADD ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2
ADD libc.so.6  /lib64/libc.so.6 

ADD gse /
CMD ["/gse"]
```

Build image (result: 6.4Mb)
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
You can change this behaviou by using the env vars GSE_BASEPATH and GSE_PORT

