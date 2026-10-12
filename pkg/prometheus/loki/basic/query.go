package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) Query(q string) (*query_result.Result, error) {
	r := query.New()

	if e := c.Get(
		c.base.Copy().Path(constant.LokiQuery).SetInteger64(
			webConstant.ParameterTime,
			time.Now().Unix(),
		).Set(webConstant.ParameterQuery, q).String(),
		&r,
	); e != nil {
		return nil, e
	}

	if r.Status != constant.Success {
		return nil, unexpected.Format("loki status: %s", r.Status)
	}

	return r.Result, nil
}
