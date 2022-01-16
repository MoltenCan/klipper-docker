# setup env to make it easy to run
# make sure you source this file

function printbox-up {
    docker run -it \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /dev:/dev \
        -v $(pwd)/test:/printbox \
        --privileged \
        moltencan/printbox \
        up
}

function printbox-listports {
    docker run -it \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /dev:/dev \
        --privileged \
        moltencan/printbox \
        list
}

echo "actions available:"
echo " printbox-up"
echo " printbox-listports"
