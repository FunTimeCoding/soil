package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/request"
)

func (c *Client) Navigate(
	context string,
	locator string,
) error {
	_, e := c.send(
		constant.LadybirdNavigateMethod,
		request.NewNavigate(context, locator),
	)

	return e
}
