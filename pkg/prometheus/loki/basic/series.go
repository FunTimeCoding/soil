package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Client) Series(series string) (string, error) {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)

	return c.Text(
		c.base.Copy().Path(constant.LokiSeries).SetInteger64(
			webConstant.ParameterStart,
			oneWeekAgo.Unix(),
		).SetInteger64(
			webConstant.ParameterEnd,
			now.Unix(),
		).Set("match[]", series).String(),
	)
}
