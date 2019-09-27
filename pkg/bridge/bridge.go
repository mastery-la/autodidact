// Package bridge defines common interfaces used for bridging a node with controllers
// to carry out real world side effects as a result of component methods
package bridge

import (
	"github.com/mastery-la/autodidact/pkg/controller"
	"github.com/mastery-la/autodidact/pkg/node"
)

type Bridge struct {
	Node        *node.Node
	Controllers []*controller.Controller
}
