#!/bin/sh
set -xe

go build main.go
cd lkm
make

sudo insmod hello.ko
sleep 2
sudo rmmod hello
sudo cat /tmp/go-crash.log
