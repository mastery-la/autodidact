package component

import (
	"github.com/mastery-la/autodidact/pkg/attribute"
)

const TypeThermostat = "ThermostatComponent"

type Thermostat struct {
	*Component
	CurrentTemperature *attribute.Temperature
	TargetTemperature  *attribute.Temperature
}

func NewThermostat(id string) *Thermostat {
	ts := new(Thermostat)

	ts.Component = New(id, TypeThermostat)

	ts.CurrentTemperature = attribute.NewTemperature()
	ts.AddAttribute(ts.CurrentTemperature.Attribute)

	ts.TargetTemperature = attribute.NewTemperature()
	ts.AddAttribute(ts.TargetTemperature.Attribute)

	return ts
}
