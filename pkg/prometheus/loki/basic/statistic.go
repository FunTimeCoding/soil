package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/web/parameter"
	"time"
)

func (c *Client) Statistic(query string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		c.base.Copy().Path(constant.Statistic).SetInteger64(
			parameter.Start,
			oneWeekAgo.UnixNano(),
		).SetInteger64(
			parameter.End,
			now.UnixNano(),
		).Set(parameter.Query, query).String(),
	)

	return result
}
