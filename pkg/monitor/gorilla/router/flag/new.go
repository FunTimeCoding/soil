package flag

import "github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"

func New(
	identifier string,
	c *client.Client,
) *Flag {
	return &Flag{Identifier: identifier, By: []*client.Client{c}}
}
