GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gse gse.go
strip gse
sudo docker build -t gse .
