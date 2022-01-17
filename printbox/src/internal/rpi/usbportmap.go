package rpi

var PortsPi3B13 = []*USBInfo{
	{
		Path:  "/sys/bus/usb/devices/1-1.1.2",
		Match: ":1.1.2",
		Alias: "TopLeft",
	},
	{
		Path:  "/sys/bus/usb/devices/1-1.1.3",
		Match: ":1.1.3",
		Alias: "BottomLeft",
	},
	{
		Path:  "/sys/bus/usb/devices/1-1.3",
		Match: ":1.3",
		Alias: "TopRight",
	},
	{
		Path:  "/sys/bus/usb/devices/1-1.2",
		Match: ":1.2",
		Alias: "BottomRight",
	},
}

var PortsPi4B = []*USBInfo{
	{
		// platform-fd500000.pcie-pci-0000:01:00.0-usb-0:1.3:1.0
		Path:  "/sys/bus/usb/devices/1-1.3",
		Match: ":1.3",
		Alias: "TopLeft",
	},
	{
		// platform-fd500000.pcie-pci-0000:01:00.0-usb-0:1.4:1.0
		Path:  "/sys/bus/usb/devices/1-1.4",
		Match: ":1.4",
		Alias: "BottomLeft",
	},
	{
		// platform-fd500000.pcie-pci-0000:01:00.0-usb-0:1.1:1.0
		Path:  "/sys/bus/usb/devices/1-1.1",
		Match: ":1.1",
		Alias: "TopRight",
	},
	{
		// platform-fd500000.pcie-pci-0000:01:00.0-usb-0:1.2:1.0
		Path:  "/sys/bus/usb/devices/1-1.2/",
		Match: ":1.2",
		Alias: "BottomRight",
	},
}
