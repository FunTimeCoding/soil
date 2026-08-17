package opnsense

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"github.com/funtimecoding/soil/pkg/opnsense/search"
)

func searchRows[T any](
	c *Client,
	path string,
	phrase string,
) ([]T, error) {
	b, e := c.basic.Post(path, search.New(phrase))

	if e != nil {
		return nil, e
	}

	var out response.Rows[T]

	if f := json.Unmarshal(b, &out); f != nil {
		return nil, f
	}

	return out.Rows, nil
}
