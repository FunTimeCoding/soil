package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/request"
)

func (c *Client) Navigate(
	context string,
	locator string,
) error {
	_, e := c.send(
		constant.NavigateMethod,
		request.NewNavigate(context, locator),
	)

	return e
}
