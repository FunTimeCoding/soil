package loki

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"time"
)

func (c *Client) MustQueryRange(
	query string,
	start time.Time,
	end time.Time,
	limit int,
) ([]*message.Message, *message.Meta) {
	result, meta, e := c.QueryRange(query, start, end, limit)
	errors.PanicOnError(e)

	return result, meta
}
