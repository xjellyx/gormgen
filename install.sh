#!/bin/bash
dir=$GOPATH
echo $dir

go build -o gormgen ./cmd/gormgen/main.go

mv gormgen $dir/bin

echo "install success"