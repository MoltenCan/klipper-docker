#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

[ -z $KDIR ] && {
    echo "KDIR missing from env"
    exit 1
}

[ -e /shared/${KDIR} ] || {
    echo "no shared folder /shared/${KDIR}"
    exit 1
}

echo "Starting Klippy"
# python3 /klipper/klippy/klippy.py -l /shared/${KDIR}/klippy.log /shared/${KDIR}/klipper.cfg
python3 /klipper/klippy/klippy.py /shared/${KDIR}/klipper.cfg
