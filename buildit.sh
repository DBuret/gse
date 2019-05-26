GGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-s -w" .
upx --brute gse
chmod 555 gse
chmod 444 template.html
sudo docker build -t gse .
sudo docker save gse > gse.tar
gzip gse.tar
ls -l gse.tar.gz
sudo docker images