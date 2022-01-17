#!/bin/sh

SRC_DIR="klipraker/src"

[ -e $SRC_DIR ] || {
    mkdir -p $SRC_DIR
}

function update() {
    echo $SRC_DIR/$1
    if [ -e $SRC_DIR/$1 ]; then
        git -C $SRC_DIR/$1 pull
    else
        git -C $SRC_DIR clone --depth 1 $2
    fi
}

function updatereq() {
    cat $SRC_DIR/klipper/scripts/klippy-requirements.txt | grep -v Jinja2 >alpine-kbase/requirements.txt
    cat $SRC_DIR/moonraker/scripts/moonraker-requirements.txt >>alpine-kbase/requirements.txt
}

case $1 in
klipper)
    update klipper https://github.com/MoltenCan/klipper.git
    updatereq
    ;;
moonraker)
    update moonraker https://github.com/Arksine/moonraker.git
    updatereq
    ;;
fluidd)
    update fluidd https://github.com/fluidd-core/fluidd.git
    ;;
all)
    $0 klipper
    $0 moonraker
    $0 fluidd
    ;;
esac
