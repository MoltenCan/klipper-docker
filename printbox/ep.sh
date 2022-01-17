#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

function check_sock {
    [ -e /var/run/docker.sock ] || {
        echo "no /var/run/docker.sock"
        exit 1
    }
}

[ -z $SHARED_PATH ] && {
    echo "Dfaulting SHARED_PATH to /shared"
    SHARED_PATH="/shared"
}

case "$1" in
up)
    check_sock
    printbox -o ${SHARED_PATH}/docker-compose.yml || exit 1
    docker-compose -f ${SHARED_PATH}/docker-compose.yml up -d
    ;;
down)
    check_sock
    docker-compose -f ${SHARED_PATH}/docker-compose.yml down
    ;;
list)
    printbox
    ;;
*)
    echo "sh | up | down | list"
    ;;
esac
