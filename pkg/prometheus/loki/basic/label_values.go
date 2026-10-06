package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) LabelValues(
	start time.Time,
	end time.Time,
	label string,
) ([]string, error) {
	r := response.NewList()

	if e := c.Get(
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
		&r,
	); e != nil {
		return nil, e
	}

	if r.Status != constant.Success {
		return nil, unexpected.Format("loki status: %s", r.Status)
	}

	return r.Labels, nil
}
