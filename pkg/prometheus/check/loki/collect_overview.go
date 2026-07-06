package loki

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/constant"
	"time"
)

func collectOverview(
	c *loki.Client,
	namespaces []string,
	since time.Duration,
) []*overview {
	end := time.Now()
	start := end.Add(-since)
	var result []*overview

	for _, n := range namespaces {
		r, _ := c.QueryRange(
			fmt.Sprintf(`{namespace="%s"}`, n),
			start,
			end,
			constant.MaximumLimit,
		)
		result = append(
			result,
			&overview{
				Namespace: n,
				Count:     len(r),
				Latest:    latest(r),
			},
		)
	}

	return result
}
