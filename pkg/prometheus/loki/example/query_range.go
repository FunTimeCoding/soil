package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
	"time"
)

func QueryRange() {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.Add(-24 * time.Hour)
	r, m := c.MustQueryRange(
		`{namespace="bot"} | json | msg="request_start", http_route="/github"`,
		start,
		end,
		constant.LokiMaximumLimit,
	)

	if false {
		console.Format("QueryRange: %s %+v\n", m.Type, m.Statistic)
	}

	for _, v := range r {
		if false {
			if v.Stream != constant.Stdout {
				continue
			}
		}

		route := v.Value(webConstant.TelemetryRoute)
		body := v.Value(webConstant.TelemetryBody)

		if strings.HasPrefix(route, "/test-") {
			continue
		}

		console.Format(
			"%s %s %s %s\n",
			v.Time.Format(timeConstant.DateMinute),
			consoleConstant.Cyan("%s", route),
			v.Stream,
			body,
		)
		h := v.ReadHeader()

		for k := range h {
			console.Format("  Header %s: %s\n", k, h.Get(k))
		}

		if v.Text != "" {
			console.Format("  Text: >%s<\n", v.Text)
		}
	}
}
