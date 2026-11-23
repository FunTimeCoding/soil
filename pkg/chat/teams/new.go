package teams

import "github.com/funtimecoding/soil/pkg/microsoft/graph"

func New(g *graph.Client) *Client {
	return &Client{graph: g}
}
