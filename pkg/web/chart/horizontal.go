package chart

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"time"
)

func (c *Chart) horizontal(at time.Time) float64 {
	span := c.end.Sub(c.start).Seconds()
	offset := at.Sub(c.start).Seconds()
	width := constant.ChartViewWidth -
		constant.ChartMarginLeft -
		constant.ChartMarginRight

	return constant.ChartMarginLeft + offset/span*width
}
