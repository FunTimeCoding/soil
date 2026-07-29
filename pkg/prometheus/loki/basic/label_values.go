package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	web "github.com/funtimecoding/soil/pkg/web/constant"
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
				constant.LokiLabel,
				label,
				constant.LokiValues,
			).SetInteger64(
				web.ParameterStart,
				start.Unix(),
			).SetInteger64(
				web.ParameterEnd,
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
