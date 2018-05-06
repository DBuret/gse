FROM scratch
ADD libpthread.so.0 /lib64/libpthread.so.0
ADD ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2
ADD libc.so.6  /lib64/libc.so.6 

ADD gse /
CMD ["/gse"]
