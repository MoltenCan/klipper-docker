FROM --platform=$TARGETPLATFORM python:alpine as builder
RUN apk add git gcc musl-dev g++ libffi-dev
RUN git clone https://github.com/MoltenCan/klipper.git
RUN pip install -r klipper/scripts/klippy-requirements.txt
COPY ep.sh /usr/local/ep.sh

ENTRYPOINT [ "/usr/local/ep.sh" ]

