package device

type teapot struct {
	name string
}

func (t teapot) GetInfo() string {
	return "I am a teapot named " + t.name
}

// ExamplePrintDeviceInfo ensures that PrintDeviceInfo prints the expected information.
// If attributes about a device are changed, calling PrintDeviceInfo after the changed
// should reflect the new values and print new information.
func ExamplePrintDeviceInfo() {
	// create an example teapot named kitchen
	exampleteapot := teapot{name: "kitchen"}
	PrintDeviceInfo(exampleteapot)

	// rename the teapot and then try printing device info
	exampleteapot.name = "break room"
	PrintDeviceInfo(exampleteapot)

	// Output:
	// I am a teapot named kitchen
	// I am a teapot named break room
}
