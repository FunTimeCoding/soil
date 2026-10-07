package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/prometheus/types/message_meta"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
	limit int,
) ([]*message.Message, *message_meta.Meta, error) {
	r, e := c.basic.QueryRange(query, start, end, limit)

	if e != nil {
		return nil, nil, e
	}

	result, meta := message.NewSlice(r)

	return result, meta, nil
}
