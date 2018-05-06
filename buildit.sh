GGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gse gse.go
strip gse
sudo docker build -t gse .
sudo docker save gse > gse.tar
gzip gse.tar
ls -l gse.tar.gz
sudo docker images
