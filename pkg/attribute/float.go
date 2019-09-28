package attribute

// FormatFloat is Float64Format
const FormatFloat = "Float64Format"

// Float represents an attribute of a float value
type Float struct {
	*Attribute
}

// NewFloat returns a new Float Attribute
func NewFloat(typ string) *Float {
	f := new(Float)

	f.Attribute = New(typ)

	f.format = FormatFloat
	f.value = 0.0

	return f
}

// SetValue sets the value for the Attribute from a float64
func (a *Float) SetValue(v float64) {
	a.UpdateValue(v)
}

// GetValue returns the vlaue for the Attribute as a float64
func (a *Float) GetValue() float64 {
	return a.Attribute.GetValue().(float64)
}
