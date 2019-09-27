package attribute

const FormatFloat = "Float64Format"

type Float struct {
	*Attribute
}

func NewFloat(typ string) *Float {
	f := new(Float)

	f.Attribute = New(typ)

	f.format = FormatFloat
	f.value = 0.0

	return f
}

func (a *Float) SetValue(v float64) {
	a.UpdateValue(v)
}

func (a *Float) GetValue() float64 {
	return a.Attribute.GetValue().(float64)
}
