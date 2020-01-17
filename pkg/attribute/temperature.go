package attribute

import (
	"encoding/json"
)

// TypeTemperature is TemperatureAttribute
const TypeTemperature = "TemperatureAttribute"

// Temperature represents an attribute of a temperature value
type Temperature struct {
	*Float
}

// NewTemperature returns a new Temperature Attribute
func NewTemperature(id string) *Temperature {
	temp := new(Temperature)

	temp.Float = NewFloat(id, TypeTemperature)

	return temp
}

// MarshalJSON implements custom JSON marshalling for temperature characteristics
func (a *Temperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.GetValue())
}
