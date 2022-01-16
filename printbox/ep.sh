#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

[ -e /var/run/docker.sock ] || {
    echo "no /var/run/docker.sock"
    exit 1
}

case "$1" in
up)
    printbox || exit 1
    docker-compose -f /printbox/docker-compose.yml up -d
    ;;
down)
    docker-compose -f /printbox/docker-compose.yml down
    ;;
*)
    echo "sh | up | down"
    ;;
esac
