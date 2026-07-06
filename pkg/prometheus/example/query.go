package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"time"
)

func Query() {
	c := prometheus.NewEnvironment()
	t := time.Now()

	for _, k := range maps.StringKeys(c.QueryIntegers(constant.Up, t)) {
		fmt.Printf("Up: %s\n", k)
	}

	countPerScrapeJob := c.QueryIntegers(
		`count by (job)({__name__=~".+"})`,
		t,
	)

	for _, k := range maps.StringKeys(countPerScrapeJob) {
		fmt.Printf(
			"Scrape Job: %s Count: %d\n",
			k,
			countPerScrapeJob[k],
		)
	}

	cardinalityPerMetric := c.QueryIntegers(
		`count by (__name__)({__name__=~".+"})`,
		t,
	)

	for _, k := range maps.StringKeys(cardinalityPerMetric) {
		fmt.Printf(
			"Metric: %s Count: %d\n",
			k,
			cardinalityPerMetric[k],
		)
	}

	// TODO: prometheus_tsdb_symbol_table_size_bytes
	fmt.Printf(
		"Load: %.1f %.1f %.1f\n",
		c.QueryFloat(constant.Load1, t),
		c.QueryFloat(constant.Load5, t),
		c.QueryFloat(constant.Load15, t),
	)

	for _, r := range parse.Generic(c.MustQuery(constant.Load1, t).Value) {
		fmt.Printf("  %s %s %s\n", r.Metric, r.Time, r.Value)
	}
}
