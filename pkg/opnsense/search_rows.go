package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"github.com/funtimecoding/soil/pkg/opnsense/search"
)

func searchRows[T any](
	c *Client,
	path string,
	phrase string,
) ([]T, error) {
	var out response.Rows[T]

	if e := c.basic.Post(path, search.New(phrase), &out); e != nil {
		return nil, e
	}

	return out.Rows, nil
}
