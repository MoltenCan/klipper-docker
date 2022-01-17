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
