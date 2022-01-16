FROM --platform=$TARGETPLATFORM moltencan/alpine-kbase
COPY klipper/ klipper/
COPY klipper_default.cfg klipper/klipper_default.cfg

COPY ep.sh /ep.sh
ENTRYPOINT [ "/ep.sh" ]

