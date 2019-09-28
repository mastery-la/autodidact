// Package attribute implements semantic values that can be read from or written to
package attribute

// ChangeFunc is a function that is called when the value changes
type ChangeFunc func(a *Attribute, newValue, oldValue interface{})

// Attribute represents a value of a Component and methods for changing
// the internal value in a safe way, also notifying listerns of changes
type Attribute struct {
	attributeType string
	format        string
	value         interface{}

	onChange *ChangeFunc
}

// New returns an Attribute for a provided type
func New(typ string) *Attribute {
	a := new(Attribute)

	a.attributeType = typ
	a.value = nil

	return a
}

// GetType returns the type of the Attribute
func (a *Attribute) GetType() string {
	return a.attributeType
}

// GetFormat returns the format of the Attribute
func (a *Attribute) GetFormat() string {
	return a.format
}

// UpdateValue updates the value of the Attribute
func (a *Attribute) UpdateValue(v interface{}) {
	a.updateValue(v)
}

// GetValue gets the value of the Attribute as a genric interface{}
func (a *Attribute) GetValue() interface{} {
	return a.value
}

// OnChange registers a callback ChangeFunc for when the value is changed
func (a *Attribute) OnChange(changeFunc ChangeFunc) {
	a.onChange = &changeFunc
}

func (a *Attribute) updateValue(value interface{}) {
	if a.onChange != nil {
		onChange := *a.onChange
		onChange(a, value, a.value)
	}
	a.value = value
}
