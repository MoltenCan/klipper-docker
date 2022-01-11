#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

case "$1" in
up)
    printbox || exit 1
    docker-compse -f /printbox/docker-compose.yml up -d
    ;;
down)
    docker-compse -f /printbox/docker-compose.yml down
    ;;
*)
    echo "sh | up | down"
    ;;
esac
