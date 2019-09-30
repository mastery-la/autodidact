package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/mastery-la/autodidact/pkg/node"
)

// Thermostat is a wrapper for a Thermostat Node to add Accessoriable support
type Thermostat struct {
	*node.Thermostat
}

// NewThermostat wraps the ThermostatNode
func NewThermostat(n *node.Thermostat) *Thermostat {
	ts := new(Thermostat)
	ts.Thermostat = n
	return ts
}

// GetAccessory returns an Accessory for the Thermostat
func (p *Thermostat) GetAccessory() *accessory.Accessory {
	ts := accessory.NewThermostat(GetInfo(p.Node), 22.0, 15.0, 32.0, 0.1)

	// ======
	// CurrentTemperature is read only (remote can GET)
	// ======

	// when remote GET CurrentTemperature, return from Thermostat Node
	ts.Thermostat.CurrentTemperature.OnValueRemoteGet(func() float64 {
		return p.CurrentTemperature.GetValue()
	})

	// SET initial value for CurrentTemperature
	ts.Thermostat.CurrentTemperature.SetValue(p.CurrentTemperature.GetValue())

	// when Thermostat Node has a new value, SET CurrentTemperature
	p.CurrentTemperature.OnNewValue(func(v interface{}) {
		ts.Thermostat.CurrentTemperature.SetValue(v.(float64))
	})

	// ======
	// TargetTemperature is read/write (remote can GET/SET)
	// ======

	// when remote GET TargetTemperature, return from Thermostat Node
	ts.Thermostat.TargetTemperature.OnValueRemoteGet(func() float64 {
		return p.TargetTemperature.GetValue()
	})

	// when remote SET TargetTemperature, set with Thermostat Node
	ts.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) {
		p.TargetTemperature.SetValue(temp)
	})

	// SET initial value for TargetTemperature
	ts.Thermostat.TargetTemperature.SetValue(p.TargetTemperature.GetValue())

	// when Thermostat Node has a new value, SET TargetTemperature
	p.TargetTemperature.OnNewValue(func(v interface{}) {
		ts.Thermostat.TargetTemperature.SetValue(v.(float64))
	})

	return ts.Accessory
}
