#!/bin/sh

echo "leeloo dallas multibuild"

platforms="linux/arm/v7,linux/arm/v6,linux/amd64,linux/arm64/v8"
ver="0.0.1"

function buildit {
    push=""
    [ "${1}" = "push" ] && {
        push="--push --tag moltencan/${tag}:latest"
    }

    cd $tag
    docker buildx build \
        --platform "${platforms}" \
        --tag moltencan/${tag}:${ver} \
        $push \
        .
}

echo "if you get an error do"
echo docker buildx create --use

case $1 in
all)
    tag="printbox"
    buildit $2
    ;;
printbox)
    tag="printbox"
    buildit $2
    ;;
klipper)
    tag="klipper"
    buildit $2
    ;;
fluidd)
    tag="fluidd"
    buildit $2
    ;;
alpine-kbase)
    tag="$1"
    buildit $2
    ;;
*)
    echo "multibuild <target> [push]"
    echo "targets  all | printbox | klipper | fluidd"
    ;;
esac
