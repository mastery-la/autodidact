package node

import (
	"encoding/json"

	"github.com/mastery-la/autodidact/pkg/component"
)

// TypeThermostat is ThermostatNode
const TypeThermostat = "ThermostatNode"

// Thermostat is a type of Node that represents a thermostat
type Thermostat struct {
	*Node
	Thermostat *component.Thermostat
}

// NewThermostat creates a Thermostat Node with the provided id
func NewThermostat(id string) *Thermostat {
	ts := new(Thermostat)

	ts.Node = New(id, TypeThermostat)

	ts.Thermostat = component.NewThermostat(ts.id)
	ts.AddComponent(ts.Thermostat.Component)

	return ts
}

// SetTargetTemperature sets the target temperature set point for the thermostat
func (t *Thermostat) SetTargetTemperature(temp float64) {
	t.Thermostat.TargetTemperature.SetValue(temp)
}

// GetTargetTemperature gets the target temperature set point for the thermostat
func (t *Thermostat) GetTargetTemperature() float64 {
	return t.Thermostat.TargetTemperature.GetValue()
}

// GetCurrentTemperature gets the current temperature as measured by the thermostat
func (t *Thermostat) GetCurrentTemperature() float64 {
	return t.Thermostat.CurrentTemperature.GetValue()
}

type thermostatPayload struct {
	Node       *Node                       `json:"node"`
	Components thermostatComponentsPayload `json:"components"`
}

type thermostatComponentsPayload struct {
	Thermostat *component.Thermostat `json:"thermoastat"`
}

// MarshalJSON uses custom payload to marshal the Thermostat Node into json
func (t *Thermostat) MarshalJSON() ([]byte, error) {
	p := thermostatPayload{
		Node: t.Node,
		Components: thermostatComponentsPayload{
			Thermostat: t.Thermostat,
		},
	}

	return json.Marshal(p)
}
