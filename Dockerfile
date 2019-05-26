FROM scratch

LABEL version="4.0"
LABEL link="https://github.com/DBuret/gse"
LABEL description="Go Show Env - micro HTTP service to help understanding container orchestrators environment"

ADD gse /
ADD template.html /
CMD ["/gse"]
