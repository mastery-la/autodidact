package node

import (
	"github.com/mastery-la/autodidact/pkg/component"
)

const TypeThermostat = "ThermostatNode"

type Thermostat struct {
	*Node
	Thermostat *component.Thermostat
}

func NewThermostat(id string) *Thermostat {
	ts := new(Thermostat)

	ts.Node = New(id, TypeThermostat)

	ts.Thermostat = component.NewThermostat(ts.id)
	ts.AddComponent(ts.Thermostat.Component)

	return ts
}

func (t Thermostat) SetTemperature(temp float64) {

}

func (t Thermostat) GetTemperature() float64 {
	return 72.0
}
