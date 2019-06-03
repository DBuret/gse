
# Description

GSE is a small standalone web app designed to help debugging the environment of orchestred containers, through a very small go app.

It can be be easly packaged into a < 3 MB docker image that will run with less than 20 Mo of ram, making it easly to deploy.

The docker image is availble on dockerhub: https://hub.docker.com/r/davidburet/gse

More details are available in the [README.adoc](README.adoc) (the README.md is needed here since docker hub does not parse asciidoc)

# endpoints

It offers 4 endpoints

## display env & http request received

* usefull to understand what env and request your containers receive

## display its version in a single line of text

* usefull to display version while testing updates deployments (rolling, blue/green, canary)
* a "stamp" can be added to the version through env variables.

## healthcheck 

* can also be turned in ok or ko state through a simple http call
* usefull to play with your container orchestrator (kubernetes, nomad…​) routing system.

## logger

* outputs what you want on the log output
* usefull to test your that log gathering system, pattern matching alarms

