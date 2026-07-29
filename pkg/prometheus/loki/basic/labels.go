package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	web "github.com/funtimecoding/soil/pkg/web/constant"
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
			c.base.Copy().Path(constant.LokiLabels).SetInteger64(
				web.ParameterStart,
				start.Unix(),
			).SetInteger64(
				web.ParameterEnd,
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
