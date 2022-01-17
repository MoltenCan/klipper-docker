#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

[ -z $SHARED_PATH ] && {
    echo "Dfaulting SHARED_PATH to /shared"
    SHARED_PATH="/shared"
}

[ -e $SHARED_PATH ] || {
    mkdir -p $SHARED_PATH
}

case "$1" in
up)
    printbox -o ${SHARED_PATH}/docker-compose.yml || exit 1
    cat ${SHARED_PATH}/docker-compose.yml
    docker-compose -f ${SHARED_PATH}/docker-compose.yml up -d
    ;;
down)
    docker-compose -f ${SHARED_PATH}/docker-compose.yml down
    ;;
list)
    printbox -o /dev/null
    ;;
*)
    echo "sh | up | down | list"
    ;;
esac
