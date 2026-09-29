package command_context

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func (c *Context) Client() *client.ClientWithResponses {
	return c.client
}
