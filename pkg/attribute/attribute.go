// Package attribute implements semantic values that can be read from or written to
package attribute

type Attribute struct {
	attributeType string
	format        string
	value         interface{}
}

func New(typ string) *Attribute {
	a := new(Attribute)

	a.attributeType = typ
	a.value = nil

	return a
}

func (a *Attribute) GetType() string {
	return a.attributeType
}

func (a *Attribute) GetFormat() string {
	return a.format
}

func (a *Attribute) UpdateValue(v interface{}) {
	a.value = v
}

func (a *Attribute) GetValue() interface{} {
	return a.value
}
