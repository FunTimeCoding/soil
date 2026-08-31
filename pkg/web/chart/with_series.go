package chart

import "github.com/funtimecoding/soil/pkg/web/chart/series"

func (c *Chart) WithSeries(s ...*series.Series) *Chart {
	c.series = append(c.series, s...)

	return c
}
