package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	"github.com/funtimecoding/soil/pkg/web/parameter"
	"log"
	"time"
)

func (c *Client) LabelValues(
	start time.Time,
	end time.Time,
	label string,
) []string {
	r := response.NewList()
	notation.MustDecode(
		c.Get(
			c.base.Copy().Path(
				"%s/%s%s",
				constant.Label,
				label,
				constant.Values,
			).SetInteger64(
				parameter.Start,
				start.Unix(),
			).SetInteger64(
				parameter.End,
				end.Unix(),
			).String(),
		),
		&r,
		true,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Labels
}
