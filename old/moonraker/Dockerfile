FROM --platform=$TARGETPLATFORM moltencan/alpine-kbase
COPY moonraker/ moonraker/
COPY moonraker_default.cfg moonraker/moonraker_default.cfg

COPY ep.sh /ep.sh
ENTRYPOINT [ "/ep.sh" ]

