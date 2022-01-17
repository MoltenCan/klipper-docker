package rpi

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var CpuInfoFile = "/proc/cpuinfo"

type Hardware struct {
	Hardware string
	Revision string
	Serial   string
	Model    string
	USB      []*USBInfo
}

type USBInfo struct {
	Connected bool
	Path      string
	Match     string
	Alias     string
	Device    string
}

func Identify() (*Hardware, error) {
	bi := &Hardware{
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
		bi.USB = PortsPi3B13
	case "Raspberry Pi 4 Model B Rev 1.1":
		bi.USB = PortsPi4B
	default:
		return nil, fmt.Errorf("not a known Raspberry Pi: %s", bi.Model)
	}
	return bi, nil
}
