# export / import container

## On dev server

List images
```
$ docker images
```

Export image
```
$ docker save imageName | gzip > imageName.tar.gz
```


## On docker server
Locally import image
```
$ gunzip gse.tar.gz
$ docker load -i gse.tar
```

Run image:
```
$ docker run -p 28657:28657 gse`
```