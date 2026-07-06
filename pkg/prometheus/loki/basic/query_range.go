package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	"github.com/funtimecoding/soil/pkg/web/parameter"
	"log"
	"time"
)

func (c *Client) QueryRange(
	q string,
	start time.Time,
	end time.Time,
	limit int,
) *query_result.Result {
	r := query.New()
	notation.MustDecode(
		c.Get(
			c.base.Copy().Path(constant.QueryRange).SetInteger64(
				parameter.Start,
				start.UnixNano(),
			).SetInteger64(
				parameter.End,
				end.UnixNano(),
			).Set(
				parameter.Query,
				q,
			).SetInteger(
				parameter.Limit,
				limit,
			).String(),
		),
		&r,
		true,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Result
}
