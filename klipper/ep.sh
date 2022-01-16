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

[ -e ${PRINTBOX_DIR}/klipper.cfg ] || {
    echo "no config detected, copying default"
    sed "s#__PRINTBOX_DIR__#${PRINTBOX_DIR}#g" /klipper/klipper_default.cfg >${PRINTBOX_DIR}/klipper.cfg
    mkdir -p ${PRINTBOX_DIR}/gcode
}

echo "starting klippy"
python3 /klipper/klippy/klippy.py -l ${PRINTBOX_DIR}/klippy.log -a ${PRINTBOX_DIR}/klipper.sock ${PRINTBOX_DIR}/klipper.cfg
