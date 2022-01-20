#!/bin/sh

[ -z $PRINTBOX_DIR ] && {
    echo "PRINTBOX_DIR not defined"
    exit 1
}

[ -e $PRINTBOX_DIR ] || {
    echo "setting up $PRINTBOX_DIR"
    mkdir -p ${PRINTBOX_DIR}
    cp -v /opt/defaults/* ${PRINTBOX_DIR}/
}
sed -i "s#__PRINTBOX_DIR__#${PRINTBOX_DIR}#g" ${PRINTBOX_DIR}/klipper.cfg
[ -z $SERIALPORT ] || {
    sed -i "s#serial: /dev/klipperserial#serial: ${SERIALPORT}#g" ${PRINTBOX_DIR}/klipper.cfg
}
# start nginx
nginx
