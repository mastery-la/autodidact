package attribute

const TypeTemperature = "TemperatureAttribute"

type Temperature struct {
	*Float
}

func NewTemperature() *Temperature {
	temp := new(Temperature)

	temp.Float = NewFloat(TypeTemperature)

	return temp
}
