#!/bin/sh

case $1 in
klipper)
    cd $1
    if [ -e klipper ]; then
        cd klipper
        git pull
    else
        git clone --depth 1 https://github.com/MoltenCan/klipper.git
    fi
    ;;
moonraker)
    cd $1
    if [ -e moonraker ]; then
        cd moonraker
        git pull
    else
        git clone --depth 1 https://github.com/Arksine/moonraker.git
    fi
    ;;
fluidd)
    echo "cadriel/fluidd"
    # cd $1
    # if [ -e fluidd ]; then
    #     cd fluidd
    #     git pull

    # else
    #     git clone --depth 1 https://github.com/fluidd-core/fluidd.git
    # fi
    ;;
esac

cat klipper/klipper/scripts/klippy-requirements.txt | grep -v Jinja2 >alpine-kbase/requirements.txt
cat moonraker/moonraker/scripts/moonraker-requirements.txt >>alpine-kbase/requirements.txt
