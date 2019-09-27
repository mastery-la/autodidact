package device

import "fmt"

// Device represents a smart device which we can
// send messages to and query for metrics
type Device interface {
	GetInfo() string
}

// PrintDeviceInfo prints to stdout information about the provided device
func PrintDeviceInfo(d Device) {
	fmt.Println(d.GetInfo())
}
