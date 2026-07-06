package github

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github/code"
)

func (c *Client) SearchCode(
	query string,
	a ...any,
) ([]*code.Code, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, _, e := c.client.Search.Code(c.context, query, nil)

	if e != nil {
		return nil, e
	}

	return code.NewSlice(result.CodeResults), nil
}
