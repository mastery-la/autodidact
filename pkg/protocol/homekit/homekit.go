// Package homekit implements the Homekit Accessory Protocol
// to work with a bridge to enable communication
package homekit

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

type Transport struct {
	cfg         hc.Config
	t           hc.Transport
	accessories []Accessorizable
}

func New(pin string, port string, storage string) *Transport {
	tp := new(Transport)

	tp.cfg = hc.Config{Pin: "12344321", Port: "12345", StoragePath: "./.db"}

	return tp
}

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

func (tp *Transport) Stop() <-chan struct{} {
	return tp.t.Stop()
}

type Accessorizable interface {
	GetAccessory() *accessory.Accessory
}

func (tp *Transport) AddAccessory(a Accessorizable) {
	tp.accessories = append(tp.accessories, a)
}

func (tp *Transport) GetAccessories() []*accessory.Accessory {
	accessories := make([]*accessory.Accessory, len(tp.accessories))

	for i, a := range tp.accessories {
		accessories[i] = a.GetAccessory()
	}

	return accessories
}
