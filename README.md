# GSE

A small standalone web app that just display its env &amp; information about the HTTP request it received

Its goal is to be embedded in a minimal container, to help us debugging the setup of our containers orchestrator & reverse proxies.

## Compile
Goal is smallest binary, we strip symbols.
```
$ GO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o gse gse.go
$ strip gse
```

## Create docker image
The idea is to create an image from scratch.

Since glibc is not supposed to be statically linked, we have some libs to add to our image (libc, pthreads, and linker). This could be avoided with muzl but we use glibc.

Hence the Dockerfile
```
FROM scratch
ADD libpthread.so.0 /lib64/libpthread.so.0
ADD ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2
ADD libc.so.6  /lib64/libc.so.6 

ADD gse /
CMD ["/gse"]
```

```
$ sudo docker build -t gse .
```

## Use gse
GSE will by default answer on URL path = /gse and port 28657.
You can change this behaviou by using the env vars GSE_BASEPATH and GSE_PORT

