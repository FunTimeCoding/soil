package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) Series(series string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		c.base.Copy().Path(constant.LokiSeries).SetInteger64(
			web.ParameterStart,
			oneWeekAgo.Unix(),
		).SetInteger64(
			web.ParameterEnd,
			now.Unix(),
		).Set("match[]", series).String(),
	)

	return result
}
