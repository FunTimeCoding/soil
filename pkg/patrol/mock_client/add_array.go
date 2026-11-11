package mock_client

import "github.com/funtimecoding/soil/pkg/mdstat/array"

func (c *Client) AddArray(a ...*array.Array) {
	c.arrays = append(c.arrays, a...)
}
