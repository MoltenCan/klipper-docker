#!/bin/sh

echo "building"
docker build -t klipraker . && docker run \
    -e PRINTBOX_DIR=/printbox \
    -e SERIALPORT=/dev/tty.usbserial-1210 \
    -it \
    -v /dev/tty.usbserial-1210:/dev/tty.usbserial-1210 \
    --privileged \
    klipraker $@
