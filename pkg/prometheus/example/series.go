package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Series() {
	c := prometheus.NewEnvironment()
	r := c.MustSeries()
	fmt.Printf("Statistics: %+v\n", r.HeadStats)
	fmt.Printf(
		"SeriesCountByLabelValuePair: %d\n",
		len(r.SeriesCountByLabelValuePair),
	)
	fmt.Printf(
		"SeriesCountByMetricName: %d\n",
		len(r.SeriesCountByMetricName),
	)
	fmt.Printf(
		"LabelValueCountByLabelName: %d\n",
		len(r.LabelValueCountByLabelName),
	)
	fmt.Printf(
		"MemoryInBytesByLabelName: %d\n",
		len(r.MemoryInBytesByLabelName),
	)

	if true {
		return
	}

	for _, s := range r.SeriesCountByMetricName {
		fmt.Printf("SeriesCountByMetricName: %+v\n", s)
	}

	for _, s := range r.SeriesCountByLabelValuePair {
		fmt.Printf("SeriesCountByLabelValuePair: %+v\n", s)
	}
}
