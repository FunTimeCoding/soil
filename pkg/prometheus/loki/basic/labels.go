package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/soil/pkg/web/parameter"
	"log"
	"time"
)

func (c *Client) Labels(
	start time.Time,
	end time.Time,
) []string {
	r := response.NewList()
	notation.MustDecode(
		c.Get(
			c.base.Copy().Path(constant.Labels).SetInteger64(
				parameter.Start,
				start.Unix(),
			).SetInteger64(
				parameter.End,
				end.Unix(),
			).String(),
		),
		&r,
		false,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Labels
}
