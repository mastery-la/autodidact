// Package node implements smart nodes that are made up of controllable components
package node

import (
	"encoding/json"

	"github.com/mastery-la/autodidact/pkg/component"
)

// Node represents a smart node which we can
// send messages to and query for metrics
type Node struct {
	id         string
	nodeType   string
	components []*component.Component
}

// New returns a new Node given an id and type
func New(id string, typ string) *Node {
	node := new(Node)

	node.id = id
	node.nodeType = typ

	return node
}

// GetID returns the ID of the Node
func (n *Node) GetID() string {
	return n.id
}

// GetType returns the Type of the Node
func (n *Node) GetType() string {
	return n.nodeType
}

// GetComponents returns an array of all Components of the Node
func (n *Node) GetComponents() []*component.Component {
	return n.components
}

// AddComponent adds a Component to the Node
func (n *Node) AddComponent(c *component.Component) {
	n.components = append(n.components, c)
}

type nodePayload struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// MarshalJSON uses custom payload to marshal the Node into json
func (n *Node) MarshalJSON() ([]byte, error) {
	p := nodePayload{
		ID:   n.GetID(),
		Type: n.GetType(),
	}

	return json.Marshal(p)
}
