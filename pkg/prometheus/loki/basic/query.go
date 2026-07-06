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

func (c *Client) Query(q string) *query_result.Result {
	r := query.New()
	notation.MustDecode(
		c.Get(
			c.base.Copy().Path(constant.Query).SetInteger64(
				parameter.Time,
				time.Now().Unix(),
			).Set(parameter.Query, q).String(),
		),
		&r,
		false,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Result
}
