package basic

import "github.com/funtimecoding/soil/pkg/web/requester/request"

func (c *Client) Get(
	path string,
	query map[string]string,
	out any,
) error {
	q := request.Get(path)

	for k, v := range query {
		q.WithParameter(k, v)
	}

	return c.requester.Notation(q, out)
}
