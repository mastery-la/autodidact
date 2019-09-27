// Package component defines a collection of read/write attributes of a node
package component

import (
	"github.com/mastery-la/autodidact/pkg/attribute"
)

type Component struct {
	id          string
	serviceType string
	attributes  []*attribute.Attribute
}

func New(id string, serviceType string) *Component {
	c := new(Component)

	c.id = id
	c.serviceType = serviceType

	return c
}

func (c *Component) GetID() string {
	return c.id
}

func (c *Component) GetType() string {
	return c.serviceType
}

func (c *Component) GetAttributes() []*attribute.Attribute {
	return c.attributes
}

func (c *Component) AddAttribute(a *attribute.Attribute) {
	c.attributes = append(c.attributes, a)
}
