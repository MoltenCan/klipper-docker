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

case "$1" in
up)
    check_sock
    printbox || exit 1
    docker-compose -f /printbox/docker-compose.yml up -d
    ;;
down)
    check_sock
    docker-compose -f /printbox/docker-compose.yml down
    ;;
list)
    printbox
    ;;
*)
    echo "sh | up | down | list"
    ;;
esac
