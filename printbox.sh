# setup env to make it easy to run
# make sure you source this file

function printbox {
    docker run -it \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /dev:/dev \
        -v printbox:/printbox \
        --privileged \
        moltencan/printbox \
        $@
}

function printbox_compile {
    [ -z $1 ] && {
        echo "please specify the port e.g. /dev/ttyAMA0"
        return
    }
    docker run -it \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v /printbox:/printbox \
        --device ${1}:/dev/klipperserial \
        --workdir "/klipper" \
        -p 8000:8000 \
        moltencan/klipper \
        compile
}

function printbox_pull {
    docker pull moltencan/printbox
    docker pull moltencan/klipper
    docker pull moltencan/moonraker
    docker pull cadriel/fluidd
}

case $1 in
pull)
    printbox_pull
    ;;
compiler)
    printbox_compile $@
    ;;
start)
    printbox up -d
    ;;
stop)
    printbox down
    ;;
*)
    echo "pull | compiler | start | stop"
    ;;

esac
