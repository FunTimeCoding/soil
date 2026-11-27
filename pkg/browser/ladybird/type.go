package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/request"
)

func (c *Client) Type(
	context string,
	text string,
) error {
	if _, e := c.send(
		constant.LadybirdActionsMethod,
		request.NewTyping(context, text),
	); e != nil {
		return e
	}

	_, e := c.send(constant.LadybirdReleaseMethod, request.NewTarget(context))

	return e
}
