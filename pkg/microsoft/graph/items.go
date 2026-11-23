package graph

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/microsoft/graph/response"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/url"
)

func (c *Client) Items(
	path string,
	v url.Values,
	limit int,
) ([]json.RawMessage, error) {
	var result []json.RawMessage
	q := request.Get(path).WithParameters(v)

	for {
		p := response.NewPage()

		if e := c.requester.Notation(q, p); e != nil {
			return nil, e
		}

		result = append(result, p.Value...)

		if limit > 0 && len(result) >= limit {
			return result[:limit], nil
		}

		if p.NextLink == "" {
			return result, nil
		}

		q = request.Absolute(p.NextLink)
	}
}
