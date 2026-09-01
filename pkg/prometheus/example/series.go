package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Series() {
	c := prometheus.NewEnvironment()
	r := c.MustSeries()
	console.Format("Statistics: %+v\n", r.HeadStats)
	console.Format(
		"SeriesCountByLabelValuePair: %d\n",
		len(r.SeriesCountByLabelValuePair),
	)
	console.Format(
		"SeriesCountByMetricName: %d\n",
		len(r.SeriesCountByMetricName),
	)
	console.Format(
		"LabelValueCountByLabelName: %d\n",
		len(r.LabelValueCountByLabelName),
	)
	console.Format(
		"MemoryInBytesByLabelName: %d\n",
		len(r.MemoryInBytesByLabelName),
	)

	if true {
		return
	}

	for _, s := range r.SeriesCountByMetricName {
		console.Format("SeriesCountByMetricName: %+v\n", s)
	}

	for _, s := range r.SeriesCountByLabelValuePair {
		console.Format("SeriesCountByLabelValuePair: %+v\n", s)
	}
}
