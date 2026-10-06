package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"
)

func (c *Client) Push(s ...*stream.Stream) error {
	return c.Post(constant.LokiPush, notation.Marshal(stream.NewPayload(s...)))
}
