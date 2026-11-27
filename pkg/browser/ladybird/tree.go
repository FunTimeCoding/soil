package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"
)

func (c *Client) Tree() ([]*response.Context, error) {
	t, e := decodeResult[response.Tree](
		c.send(constant.LadybirdTreeMethod, map[string]any{}),
	)

	if e != nil {
		return nil, e
	}

	return t.Contexts, nil
}
