#!/bin/sh

echo "leeloo dallas multibuild"

[ -z $1 ] && {
    echo "need a target"
    echo "targets  all | printbox | klipraker | alpine-kbase"
    exit 1
}

platforms="linux/arm/v7,linux/arm/v6,linux/amd64,linux/arm64/v8"
ver="0.0.1"

function buildit {
    [ -e $1 ] || {
        echo "$1 doesn't exist"
        exit 1
    }
    cd $1
    docker buildx build \
        --platform "${platforms}" \
        --tag moltencan/${1}:latest \
        --push \
        .
}

echo "if you get an error do"
echo docker buildx create --use

case $1 in
all)
    $0 "alpine-kbase"
    $0 "printbox"
    $0 "klipraker"
    ;;
*)
    buildit $1
    ;;
esac
