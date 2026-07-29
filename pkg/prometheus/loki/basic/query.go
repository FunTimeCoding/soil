package basic

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"log"
	"time"
)

func (c *Client) Query(q string) *query_result.Result {
	r := query.New()
	notation.MustDecode(
		c.Get(
			c.base.Copy().Path(constant.LokiQuery).SetInteger64(
				web.ParameterTime,
				time.Now().Unix(),
			).Set(web.ParameterQuery, q).String(),
		),
		&r,
		false,
	)

	if r.Status != constant.Success {
		log.Panicf("unexpected status: %s", r.Status)
	}

	return r.Result
}
