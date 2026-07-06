package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
	limit int,
) ([]*message.Message, *message.Meta) {
	return message.NewSlice(c.basic.QueryRange(query, start, end, limit))
}
