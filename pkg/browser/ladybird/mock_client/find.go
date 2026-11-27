package mock_client

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) find(identifier string) (*response.Context, error) {
	for _, v := range c.contexts {
		if v.Identifier == identifier {
			return v, nil
		}
	}

	return nil, not_found.Format(
		"%s: %s",
		constant.LadybirdNoSuchFrame,
		identifier,
	)
}
