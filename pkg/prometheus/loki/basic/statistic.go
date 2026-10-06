package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) Statistic(query string) (string, error) {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)

	return c.Text(
		c.base.Copy().Path(constant.LokiStatistic).SetInteger64(
			web.ParameterStart,
			oneWeekAgo.UnixNano(),
		).SetInteger64(
			web.ParameterEnd,
			now.UnixNano(),
		).Set(web.ParameterQuery, query).String(),
	)
}
