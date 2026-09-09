package face

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

type MetricSource interface {
	Query(
		q string,
		t time.Time,
	) (*query_result.Result, error)
}
