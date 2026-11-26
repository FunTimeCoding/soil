package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/request"
)

func (c *Client) Subscribe(events ...string) error {
	_, e := c.send(constant.SubscribeMethod, request.NewSubscribe(events))

	return e
}
