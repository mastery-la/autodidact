package component

import (
	"github.com/mastery-la/autodidact/pkg/attribute"
)

// TypeSwitch is SwitchComponent
const TypeSwitch = "SwitchComponent"

// Switch is a type of Component that has an on/off attribute
type Switch struct {
	*Component
	On *attribute.Temperature
}

// NewSwitch creates a Switch Component with a provided id
func NewSwitch(id string) *Switch {
	s := new(Switch)

	s.Component = New(id, TypeSwitch)

	s.On = attribute.NewTemperature("On")
	s.AddAttribute(s.On.Attribute)

	return s
}
