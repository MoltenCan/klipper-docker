FROM --platform=$TARGETPLATFORM python:alpine as builder
RUN apk add unzip git python3 py3-pip python3-dev \
    patch gcc musl-dev zlib-dev libjpeg-turbo-dev \
    g++ libffi-dev curl-dev bash iproute2 libsodium
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

# needed for klipper compiling
RUN apk add make python2 newlib-arm-none-eabi gcc-avr avrdude avr-libc
