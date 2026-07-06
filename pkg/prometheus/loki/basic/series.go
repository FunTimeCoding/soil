package basic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"github.com/funtimecoding/soil/pkg/web/parameter"
	"time"
)

func (c *Client) Series(series string) string {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	result := c.Get(
		c.base.Copy().Path(constant.Series).SetInteger64(
			parameter.Start,
			oneWeekAgo.Unix(),
		).SetInteger64(
			parameter.End,
			now.Unix(),
		).Set("match[]", series).String(),
	)

	return result
}
