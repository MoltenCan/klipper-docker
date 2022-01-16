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
    $0 "alpine-kbase" $2
    $0 "printbox" $2
    $0 "klipper" $2
    $0 "moonraker" $2
    $0 "fluidd" $2
    ;;
printbox)
    tag="$1"
    buildit $2
    ;;
klipper)
    tag="$1"
    buildit $2
    ;;
moonraker)
    tag="$1"
    buildit $2
    ;;
fluidd)
    echo "not using official fluidd continaer?"
    # tag="$1"
    # buildit $2
    ;;
alpine-kbase)
    tag="$1"
    buildit $2
    ;;
*)
    echo "multibuild <target> [push]"
    echo "targets  all | printbox | klipper | moonraker"
    ;;
esac
