GGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o gse gse.go
sudo docker build -t gse .
