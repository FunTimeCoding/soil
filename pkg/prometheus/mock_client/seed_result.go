package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/query_result"

func (c *Client) SeedResult(
	q string,
	r *query_result.Result,
) {
	c.results[q] = r
}
