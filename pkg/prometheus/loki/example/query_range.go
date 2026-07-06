package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	telemetry "github.com/funtimecoding/soil/pkg/web/telemetry/constant"
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
		constant.MaximumLimit,
	)

	if false {
		fmt.Printf("QueryRange: %s %+v\n", m.Type, m.Statistic)
	}

	for _, v := range r {
		if false {
			if v.Stream != constant.Stdout {
				continue
			}
		}

		route := v.Value(telemetry.Route)
		body := v.Value(telemetry.Body)

		if strings.HasPrefix(route, "/test-") {
			continue
		}

		fmt.Printf(
			"%s %s %s %s\n",
			v.Time.Format(timeLibrary.DateMinute),
			console.Cyan("%s", route),
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
