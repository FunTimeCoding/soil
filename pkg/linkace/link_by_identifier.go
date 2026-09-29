package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) LinkByIdentifier(identifier int) (*link.Link, error) {
	var r response.Link

	if e := c.basic.Get(
		fmt.Sprintf("links/%d", identifier),
		nil,
		&r,
	); e != nil {
		return nil, e
	}

	return link.New(r, c.host), nil
}
