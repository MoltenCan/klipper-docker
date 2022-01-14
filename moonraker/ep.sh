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

echo "starting moonraker"
python3 /moonraker/moonraker/moonraker.py /shared/${KDIR}/moonraker.cfg
