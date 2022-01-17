package printbox

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"printbox/internal/rpi"
	"strings"
)

const (
	VERSION = "0.0.4"
)

var (
	Debug      = false
	USBPath    = "/dev/serial/by-path/"
	SharedPath = "/shared"
)

func GetPorts(bi *rpi.Hardware) (map[string]string, error) {
	portMap := map[string]string{}

	fList, err := ioutil.ReadDir(USBPath)
	if err != nil {
		return portMap, err
	}
	for _, f := range fList {
		orig, err := filepath.EvalSymlinks(filepath.Join(USBPath, f.Name()))
		if err != nil {
			return portMap, err
		}
		port := orig
		alias := filepath.Base(orig)
		if bi != nil {
			for _, usb := range bi.USB {
				if strings.Contains(f.Name(), usb.Match) {
					alias = usb.Alias
				}
			}
		}
		portMap[alias] = port
	}
	return portMap, nil
}

func Debugf(f string, v ...interface{}) {
	if !Debug {
		return
	}
	msg := fmt.Sprintf(f, v...)
	fmt.Println(msg)
}

func Logf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	fmt.Println(msg)
}
