package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/postgres/types/statement"
)

func (c *Client) Query(
	_ context.Context,
	sql string,
	a ...any,
) (*result.Result, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.statements = append(c.statements, statement.New(sql, a))

	if c.Failure != nil {
		return nil, c.Failure
	}

	if v, okay := c.results[sql]; okay {
		return v, nil
	}

	return result.New(nil, nil), nil
}
