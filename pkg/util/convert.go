package util

func Fahrenheit2Celsius(temp float64) float64 {
	return (temp - 32.0) * (5.0 / 9.0)
}

func Celsius2Fahrenheit(temp float64) float64 {
	return (9.0/5.0)*temp + 32.0
}
