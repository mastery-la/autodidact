// Package homekit implements the Homekit Accessory Protocol
// to work with a bridge to enable communication
package homekit

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

// Transport represents a communication transport for the Homekit Protocol
type Transport struct {
	cfg         hc.Config
	t           hc.Transport
	accessories []Accessorizable
}

// New creates a new Transport with a given pin, port, and storage location
func New(pin string, port string, storage string) *Transport {
	tp := new(Transport)

	tp.cfg = hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./.db"}

	return tp
}

// Start starts the transport server
func (tp *Transport) Start() {
	accessories := tp.GetAccessories()

	ipt, err := hc.NewIPTransport(tp.cfg, accessories[0], accessories[1:]...)
	if err != nil {
		log.Panic(err)
	}
	tp.t = ipt

	hc.OnTermination(func() {
		<-tp.Stop()
	})

	tp.t.Start()
}

// Stop stops the transport server
func (tp *Transport) Stop() <-chan struct{} {
	return tp.t.Stop()
}

// Accessorizable is an interface that specifies a method for getting an accessory
type Accessorizable interface {
	GetAccessory() *accessory.Accessory
}

// AddAccessory appends the provided Accessorizable type to the Transport
func (tp *Transport) AddAccessory(a Accessorizable) {
	tp.accessories = append(tp.accessories, a)
}

// GetAccessories returns an array of Accessories from the list of Accessorizable types
func (tp *Transport) GetAccessories() []*accessory.Accessory {
	accessories := make([]*accessory.Accessory, len(tp.accessories))

	for i, a := range tp.accessories {
		accessories[i] = a.GetAccessory()
	}

	return accessories
}
