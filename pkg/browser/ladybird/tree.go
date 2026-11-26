package ladybird

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/response"
)

func (c *Client) Tree() ([]*response.Context, error) {
	t, e := decodeResult[response.Tree](
		c.send(constant.TreeMethod, map[string]any{}),
	)

	if e != nil {
		return nil, e
	}

	return t.Contexts, nil
}
