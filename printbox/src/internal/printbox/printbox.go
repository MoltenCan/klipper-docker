package printbox

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	VERSION = "0.0.1"
)

var (
	Debug       = false
	CpuInfoFile = "/proc/cpuinfo"
	USBPath     = "/dev/serial/by-path/"
	SharedPath  = "/printbox"
)

type BoardInfo struct {
	Hardware string
	Revision string
	Serial   string
	Model    string
	USB      []*USBInfo
	USBCount int
}

type USBInfo struct {
	Connected bool
	Path      string
	Match     string
	Position  string
	Device    string
}

func GetBoard() (*BoardInfo, error) {
	bi := &BoardInfo{
		Hardware: "unknown",
		Revision: "unknown",
		Serial:   "unknown",
		Model:    "unknown",
	}

	data, err := ioutil.ReadFile(CpuInfoFile)
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		switch {
		case strings.HasPrefix(parts[0], "Hardware"):
			bi.Hardware = parts[1]
		case strings.HasPrefix(parts[0], "Revision"):
			bi.Revision = parts[1]
		case strings.HasPrefix(parts[0], "Serial"):
			bi.Serial = parts[1]
		case strings.HasPrefix(parts[0], "Model"):
			bi.Model = parts[1]
		}
	}

	// now match the port map
	switch bi.Model {
	case "Raspberry Pi 3 Model B Plus Rev 1.3":
		Debugf("board is %s", bi.Model)
		bi.USB = PortsPi3B13
	}

	return bi, nil
}

func CheckPorts(bi *BoardInfo) error {
	bi.USBCount = 0
	fList, err := ioutil.ReadDir(USBPath)
	if err != nil {
		return err
	}
	for _, f := range fList {

		for i, usb := range bi.USB {
			if strings.Contains(f.Name(), usb.Match) {
				orig, err := filepath.EvalSymlinks(filepath.Join(USBPath, f.Name()))
				if err != nil {
					continue
				}
				usb.Device = orig
				usb.Connected = true
				bi.USBCount++

				Logf("found %s", f.Name())
				Logf(" port %d (%s)", i, usb.Position)
				Logf(" dev %s", usb.Device)
			}
		}
	}
	return nil
}

func ChkMkDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return os.MkdirAll(path, 0644)
	}
	return nil
}

func CheckDirs(bi *BoardInfo) error {
	for i, usb := range bi.USB {
		if !usb.Connected {
			continue
		}
		if err := ChkMkDir(filepath.Join(SharedPath, strconv.Itoa(i))); err != nil {
			return err
		}
	}
	return nil
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
