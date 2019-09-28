// Package bridge defines common interfaces used for bridging a node with controllers
// to carry out real world side effects as a result of component methods
package bridge

import (
	"github.com/mastery-la/autodidact/pkg/controller"
	"github.com/mastery-la/autodidact/pkg/node"
)

type Bridge struct {
	node        *node.Node
	controllers []*controller.Controller
}

func New(node *node.Node) *Bridge {
	b := new(Bridge)

	b.node = node
	b.controllers = make([]*controller.Controller, 0)

	return b
}

func (b *Bridge) GetNode() *node.Node {
	return b.node
}

func (b *Bridge) GetControllers() []*controller.Controller {
	return b.controllers
}

func (b *Bridge) AddController(c *controller.Controller) {
	b.controllers = append(b.controllers, c)
}
