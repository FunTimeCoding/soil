package loki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/message"
	"github.com/funtimecoding/soil/pkg/web/telemetry/constant"
	"sort"
	"time"
)

func collect(
	c *loki.Client,
	namespace string,
	since time.Duration,
	route string,
	m string,
	exclude []string,
	limit int,
) []*message.Message {
	end := time.Now()
	start := end.Add(-since)
	query := fmt.Sprintf(`{namespace="%s"} | json`, namespace)

	if route != "" {
		query = fmt.Sprintf(`%s | http_route="%s"`, query, route)
	}

	if m != "" {
		query = fmt.Sprintf(`%s | msg="%s"`, query, m)
	}

	if m == "" {
		for _, e := range exclude {
			query = fmt.Sprintf(`%s | msg!="%s"`, query, e)
		}
	}

	r, _ := c.QueryRange(query, start, end, limit)
	var result []*message.Message

	for _, v := range r {
		if route != "" && v.Value(constant.Route) == "" {
			continue
		}

		result = append(result, v)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Time.Before(result[j].Time)
		},
	)

	return result
}
