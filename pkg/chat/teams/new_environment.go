package teams

import "github.com/funtimecoding/soil/pkg/microsoft/graph"

func NewEnvironment() *Client {
	return New(graph.NewEnvironment())
}
