#!/bin/bash

proj="sample_web"
src="./"
dest="./build"
rm -rf $dest
mkdir -p $dest/scripts/
mkdir -p $dest/templates/
cp ./scripts/* ./build/scripts/
cp ./templates/* ./build/templates/
GOOS=linux GOARCH=amd64 go build -o $dest/$proj ./main/main.go
cur_dateTime="`date +%Y%m%d%H%M`"
cd $dest && tar -czvf sample_web_$cur_dateTime.tar.gz ./*