// Package component defines a collection of read/write attributes of a node
package component

import (
	"github.com/mastery-la/autodidact/pkg/attribute"
)

// Component represents a compontent of a Node which is made
// up of Attributes that can be read from and/or written to
type Component struct {
	id            string
	componentType string
	attributes    []*attribute.Attribute
}

// New returns a new Component given an id and type
func New(id string, typ string) *Component {
	c := new(Component)

	c.id = id
	c.componentType = typ

	return c
}

// GetID returns the ID of the Component
func (c *Component) GetID() string {
	return c.id
}

// GetType returns the type of the Component
func (c *Component) GetType() string {
	return c.componentType
}

// GetAttributes returns an array of all Attributes of the Component
func (c *Component) GetAttributes() []*attribute.Attribute {
	return c.attributes
}

// AddAttribute adds an Attribute to the Component
func (c *Component) AddAttribute(a *attribute.Attribute) {
	c.attributes = append(c.attributes, a)
}
