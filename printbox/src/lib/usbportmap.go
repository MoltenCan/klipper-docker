package lib

var PortsPi3B13 = []*USBInfo{
	{
		Path:     "/sys/bus/usb/devices/1-1.1.2",
		Match:    ":1.1.2",
		Position: "TopLeft",
	},
	{
		Path:     "/sys/bus/usb/devices/1-1.1.3",
		Match:    ":1.1.3",
		Position: "BottomLeft",
	},
	{
		Path:     "/sys/bus/usb/devices/1-1.3",
		Match:    ":1.3",
		Position: "TopRight",
	},
	{
		Path:     "/sys/bus/usb/devices/1-1.2",
		Match:    ":1.2",
		Position: "BottomRight",
	},
}
