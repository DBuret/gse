GGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-w" .
strip gse
chmod 755 gse
sudo docker build -t gse .
sudo docker images