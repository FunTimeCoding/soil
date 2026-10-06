package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) QueryRange(
	q string,
	start time.Time,
	end time.Time,
	limit int,
) (*query_result.Result, error) {
	r := query.New()

	if e := c.Get(
		c.base.Copy().Path(constant.LokiQueryRange).SetInteger64(
			web.ParameterStart,
			start.UnixNano(),
		).SetInteger64(
			web.ParameterEnd,
			end.UnixNano(),
		).Set(
			web.ParameterQuery,
			q,
		).SetInteger(
			web.ParameterLimit,
			limit,
		).String(),
		&r,
	); e != nil {
		return nil, e
	}

	if r.Status != constant.Success {
		return nil, unexpected.Format("loki status: %s", r.Status)
	}

	return r.Result, nil
}
