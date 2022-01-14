# klipper-docker
klipper components dockerised

This is probaly not what you want, I suggest you use the [official install](https://www.klipper3d.org/Installation.html) instead

components in here are based on other open source projects which are better than this one:
- https://www.klipper3d.org/
- https://github.com/Klipper3d/klipper
- https://github.com/Arksine/moonraker
- https://docs.fluidd.xyz/
- https://github.com/fluidd-core/fluidd
- https://www.alpinelinux.org/


# crazy enough to try?
so what i'm doing here is completely dockerising all the klipper components, my use case is running multiple printers off 1 single rpi  
I don't like python or messing with virtualenvs so I dockerise all the components  

once I actually getting it working you should only need to run the printbox container which will detect your serial ports, download the images and run up klippers  

why am I doing this? well primarily all the existing containers (not fluidd) don't support multi-arch (including docker-compose??) and like every other stupid dev I like to do things "my way"  

