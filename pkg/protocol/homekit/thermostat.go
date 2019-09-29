package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/mastery-la/autodidact/pkg/node"
)

type Thermostat struct {
	*node.Thermostat
}

func NewThermostat(n *node.Thermostat) *Thermostat {
	ts := new(Thermostat)
	ts.Thermostat = n
	return ts
}

func (p *Thermostat) GetAccessory() *accessory.Accessory {
	ts := accessory.NewThermostat(GetInfo(p.Node), 22.0, 15.0, 32.0, 0.1)

	ts.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) {
		p.TargetTemperature.SetValue(temp)
	})

	return ts.Accessory
}
