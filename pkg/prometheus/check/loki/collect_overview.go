package loki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"time"
)

func collectOverview(
	c *loki.Client,
	namespaces []string,
	since time.Duration,
) []*Overview {
	end := time.Now()
	start := end.Add(-since)
	var result []*Overview

	for _, n := range namespaces {
		r, _ := c.MustQueryRange(
			fmt.Sprintf(`{namespace="%s"}`, n),
			start,
			end,
			constant.LokiMaximumLimit,
		)
		result = append(
			result,
			&Overview{Namespace: n, Count: len(r), Latest: latest(r)},
		)
	}

	return result
}
