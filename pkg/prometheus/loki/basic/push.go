package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/stream"
)

func (c *Client) Push(s ...*stream.Stream) {
	c.Post(
		c.base.Copy().Path(constant.LokiPush).String(),
		notation.Marshal(stream.NewPayload(s...)),
	)
}
