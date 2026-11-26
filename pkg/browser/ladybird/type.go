package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/request"
)

func (c *Client) Type(
	context string,
	text string,
) error {
	if _, e := c.send(
		constant.ActionsMethod,
		request.NewTyping(context, text),
	); e != nil {
		return e
	}

	_, e := c.send(constant.ReleaseMethod, request.NewTarget(context))

	return e
}
