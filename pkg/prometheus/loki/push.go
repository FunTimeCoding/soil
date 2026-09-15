package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"
	"time"
)

func (c *Client) Push(
	labels map[string]string,
	lines ...string,
) {
	s := stream.New(labels)

	for _, l := range lines {
		s.Add(time.Now(), l)
	}

	c.basic.Push(s)
}
