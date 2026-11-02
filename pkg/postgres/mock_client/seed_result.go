package mock_client

import "github.com/funtimecoding/soil/pkg/postgres/types/result"

func (c *Client) SeedResult(
	sql string,
	r *result.Result,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.results[sql] = r
}
