#!/bin/sh

[ "$1" = "sh" ] && {
    /bin/sh
    exit 0
}

[ -z $PRINTBOX_DIR ] && {
    echo "PRINTBOX_DIR missing from env"
    exit 1
}

[ -e ${PRINTBOX_DIR} ] || {
    echo "no shared folder ${PRINTBOX_DIR}"
    exit 1
}

[ -e ${PRINTBOX_DIR}/moonraker.cfg ] || {
    echo "no config detected, copying default"
    cp -v /moonraker/moonraker_default.cfg ${PRINTBOX_DIR}/moonraker.cfg
}

# crate symlink to uds
ln -s ${PRINTBOX_DIR}/klipper.sock /tmp/klippy_uds

echo "starting moonraker"
python3 /moonraker/moonraker/moonraker.py -l ${PRINTBOX_DIR}/moonraker.log -c ${PRINTBOX_DIR}/moonraker.cfg
