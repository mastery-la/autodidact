// Package node implements smart nodes that are made up of controllable components
package node

import (
	"github.com/mastery-la/autodidact/pkg/component"
)

// Node represents a smart node which we can
// send messages to and query for metrics
type Node struct {
	id         string
	nodeType   string
	components []*component.Component
}

func New(id string, nodeType string) *Node {
	node := new(Node)

	node.id = id
	node.nodeType = nodeType

	return node
}

func (n *Node) GetID() string {
	return n.id
}

func (n *Node) GetType() string {
	return n.nodeType
}

func (n *Node) GetComponents() []*component.Component {
	return n.components
}

func (n *Node) AddComponent(c *component.Component) {
	n.components = append(n.components, c)
}
