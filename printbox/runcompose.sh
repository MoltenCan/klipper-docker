#!/bin/sh

docker run -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v /dev:/dev \
    -v $(pwd)/test:/printbox \
    --privileged \
    moltencan/printbox up
