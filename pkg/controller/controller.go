// Package controller defines a common interface used by various controllers
package controller

type Controller struct {
}

func New() *Controller {
	return new(Controller)
}
