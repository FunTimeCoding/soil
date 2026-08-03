package example

import (
	"fmt"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
	"time"
)

func QueryRange() {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.Add(-24 * time.Hour)
	r, m := c.QueryRange(
		`{namespace="bot"} | json | msg="request_start", http_route="/github"`,
		start,
		end,
		prometheus.LokiMaximumLimit,
	)

	if false {
		fmt.Printf("QueryRange: %s %+v\n", m.Type, m.Statistic)
	}

	for _, v := range r {
		if false {
			if v.Stream != prometheus.Stdout {
				continue
			}
		}

		route := v.Value(constant.TelemetryRoute)
		body := v.Value(constant.TelemetryBody)

		if strings.HasPrefix(route, "/test-") {
			continue
		}

		fmt.Printf(
			"%s %s %s %s\n",
			v.Time.Format(timeConstant.DateMinute),
			consoleConstant.Cyan("%s", route),
			v.Stream,
			body,
		)
		h := v.ReadHeader()

		for k := range h {
			fmt.Printf("  Header %s: %s\n", k, h.Get(k))
		}

		if v.Text != "" {
			fmt.Printf("  Text: >%s<\n", v.Text)
		}
	}
}
