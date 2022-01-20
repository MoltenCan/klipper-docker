FROM moltencan/starter:latest as starter

FROM debian:buster-slim

RUN apt-get update
RUN apt-get install -y git python3 sudo nginx procps unzip

COPY sudoers /etc/sudoers
COPY defaults/ defaults/

# fake systemd
RUN echo "#!/bin/sh\ntrue" > /bin/systemctl && chmod 755 /bin/systemctl

RUN useradd -m klipper
RUN usermod -aG sudo klipper

USER klipper
WORKDIR /home/klipper

# install klipper
RUN git clone https://github.com/Klipper3d/klipper.git
RUN ./klipper/scripts/install-debian.sh

# install moonraker
RUN git clone https://github.com/Arksine/moonraker.git
RUN ./moonraker/scripts/install-moonraker.sh

# install fluidd
RUN git clone https://github.com/fluidd-core/fluidd.git
ARG FRONTEND_ZIP_URL=https://github.com/fluidd-core/fluidd/releases/latest/download/fluidd.zip

USER root
RUN rm -f /usr/share/nginx/html/*
RUN rm -f /etc/nginx/sites-enabled/default
ADD ${FRONTEND_ZIP_URL} /tmp/frontend.zip
RUN unzip /tmp/frontend.zip -d /usr/share/nginx/html/

RUN cp /home/klipper/fluidd/server/nginx-site.conf /etc/nginx/sites-enabled/default
RUN cp -r /home/klipper/fluidd/ /usr/share/nginx/html/  
RUN cp /home/klipper/fluidd/server/config.json /usr/share/nginx/html/config.json

COPY --from=starter /bin/starter /bin/starter
COPY starter.yaml /starter.yaml
COPY starter.yaml /starter.yaml

ENTRYPOINT [ "/bin/starter" ]