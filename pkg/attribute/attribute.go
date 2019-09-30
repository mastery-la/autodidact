// Package attribute implements semantic values that can be read from or written to
package attribute

// ChangeFunc is a function that is called when the value changes
type ChangeFunc func(a *Attribute, newValue, oldValue interface{})

// NewValueFunc is a shorthand for ChangeFunc that only accepts newValue
type NewValueFunc func(newValue interface{})

// Attribute represents a value of a Component and methods for changing
// the internal value in a safe way, also notifying listerns of changes
type Attribute struct {
	id     string
	typ    string
	format string
	value  interface{}

	changeFuncs []ChangeFunc
}

// New returns an Attribute for a provided type
func New(id string, typ string) *Attribute {
	a := new(Attribute)

	a.id = id
	a.typ = typ
	a.value = nil
	a.changeFuncs = make([]ChangeFunc, 0)

	return a
}

// GetType returns the type of the Attribute
func (a *Attribute) GetType() string {
	return a.typ
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
func (a *Attribute) OnChange(fn ChangeFunc) {
	a.changeFuncs = append(a.changeFuncs, fn)
}

// OnNewValue registers a callback NewValueFunc for when the value is changed
func (a *Attribute) OnNewValue(fn NewValueFunc) {
	changeFunc := func(a *Attribute, newValue, oldValue interface{}) {
		if newValue == oldValue {
			return
		}
		fn(newValue)
	}
	a.OnChange(changeFunc)
}

func (a *Attribute) updateValue(value interface{}) {
	for _, fn := range a.changeFuncs {
		fn(a, value, a.value)
	}
	a.value = value
}
