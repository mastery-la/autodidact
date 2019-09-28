package component

import (
	"encoding/json"

	"github.com/mastery-la/autodidact/pkg/attribute"
)

// TypeThermostat is ThermosatComponent
const TypeThermostat = "ThermostatComponent"

// Thermostat is a type of Component that that has a current and target temperature
type Thermostat struct {
	*Component
	CurrentTemperature *attribute.Temperature
	TargetTemperature  *attribute.Temperature
}

// NewThermostat creates a Thermostat Component with a provided id
func NewThermostat(id string) *Thermostat {
	ts := new(Thermostat)

	ts.Component = New(id, TypeThermostat)

	ts.CurrentTemperature = attribute.NewTemperature()
	ts.AddAttribute(ts.CurrentTemperature.Attribute)

	ts.TargetTemperature = attribute.NewTemperature()
	ts.AddAttribute(ts.TargetTemperature.Attribute)

	return ts
}

type thermostatPayload struct {
	Attributes thermostatAttributesPayload `json:"attributes"`
}

type thermostatAttributesPayload struct {
	CurrentTemperature *attribute.Temperature `json:"currentTemperature"`
	TargetTemperature  *attribute.Temperature `json:"targetTemperature"`
}

// MarshalJSON uses custom payload to marshal the Thermostat Component into json
func (c *Thermostat) MarshalJSON() ([]byte, error) {
	p := thermostatPayload{
		Attributes: thermostatAttributesPayload{
			CurrentTemperature: c.CurrentTemperature,
			TargetTemperature:  c.TargetTemperature,
		},
	}

	return json.Marshal(p)
}
