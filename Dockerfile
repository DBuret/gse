FROM devalias/upx:devel AS upx

FROM golang AS builder
COPY --from=upx /usr/bin/upx /usr/bin/upx
WORKDIR /go/src/github.com/DBuret/gse
RUN go get -d -v github.com/namsral/flag
COPY gse.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-s -w" .
RUN  /usr/bin/upx --brute gse

FROM scratch
LABEL version="4.0"
LABEL link="https://github.com/DBuret/gse"
LABEL description="Go Show Env - micro HTTP service to help understanding container orchestrators environment"
WORKDIR /root/
COPY --from=builder /root/gse .
ADD template.html /
CMD ["./gse"]
