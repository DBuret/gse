FROM golang
WORKDIR /go/src/github.com/DBuret
RUN go get -d -v github.com/namsral/flag
COPY gse.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-s -w" .

FROM scratch
LABEL version="4.0"
LABEL link="https://github.com/DBuret/gse"
LABEL description="Go Show Env - micro HTTP service to help understanding container orchestrators environment"
COPY --from=0 /go/src/github.com/DBuret/gse /
ADD template.html /
CMD ["/gse"]
