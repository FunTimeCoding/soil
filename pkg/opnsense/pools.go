package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
)

func (c *Client) Pools(phrase string) ([]*pool.Pool, error) {
	rows, e := searchRows[response.Pool](c, constant.PoolSearch, phrase)

	if e != nil {
		return nil, e
	}

	return pool.NewSlice(rows), nil
}
