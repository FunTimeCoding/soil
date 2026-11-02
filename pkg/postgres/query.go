package postgres

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/postgres/refusal"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

func (c *Client) Query(
	x context.Context,
	sql string,
	a ...any,
) (*result.Result, error) {
	rows, e := c.pool.Query(x, sql, a...)

	if e != nil {
		return nil, refusal.Classify(fmt.Errorf("query: %w", e))
	}

	defer rows.Close()
	v, f := scanRows(rows)

	if f != nil {
		return nil, refusal.Classify(f)
	}

	return v, nil
}
