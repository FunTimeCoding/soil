package chart

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Chart) vertical(value float64) float64 {
	height := constant.ChartViewHeight -
		constant.ChartMarginTop -
		constant.ChartMarginBottom

	return constant.ChartViewHeight -
		constant.ChartMarginBottom -
		value*height/c.maximum
}
