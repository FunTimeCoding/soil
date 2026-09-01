package example

import (
	"github.com/funtimecoding/soil/pkg/console"
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
		console.Format("Up: %s\n", k)
	}

	countPerScrapeJob := c.QueryIntegers(`count by (job)({__name__=~".+"})`, t)

	for _, k := range maps.StringKeys(countPerScrapeJob) {
		console.Format("Scrape Job: %s Count: %d\n", k, countPerScrapeJob[k])
	}

	cardinalityPerMetric := c.QueryIntegers(
		`count by (__name__)({__name__=~".+"})`,
		t,
	)

	for _, k := range maps.StringKeys(cardinalityPerMetric) {
		console.Format("Metric: %s Count: %d\n", k, cardinalityPerMetric[k])
	}

	// TODO: prometheus_tsdb_symbol_table_size_bytes
	console.Format(
		"Load: %.1f %.1f %.1f\n",
		c.QueryFloat(constant.Load1, t),
		c.QueryFloat(constant.Load5, t),
		c.QueryFloat(constant.Load15, t),
	)

	for _, r := range parse.Generic(c.MustQuery(constant.Load1, t).Value) {
		console.Format("  %s %s %s\n", r.Metric, r.Time, r.Value)
	}
}
