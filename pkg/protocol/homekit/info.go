package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/mastery-la/autodidact/pkg/node"
)

// GetInfo returns accessory Info from a Node
func GetInfo(node *node.Node) accessory.Info {
	return accessory.Info{
		Name:             node.Name,
		SerialNumber:     node.SerialNumber,
		Manufacturer:     node.Manufacturer,
		Model:            node.Model,
		FirmwareRevision: "0.0",
	}
}
