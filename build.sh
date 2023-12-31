#!/usr/bin/env bash
export PATH="$PATH:$(go env GOPATH)/bin"

echo "  ---> make clean <---"
rm -rf output

echo "  ---> init dir <---"
mkdir output
mkdir output/conf

echo "  ---> init files <---"
cp -r conf output/
cp static/startup.sh output/

echo "  ---> make server <---"
go build -o output/server/server cmd/server/main.go

echo "  ---> make client <---"
go build -o output/client/client cmd/client/main.go

echo "  ---> make executor <---"
go build -o output/executor/executor cmd/executor/main.go