package observe_transport

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"time"
)

func (t *Transport) RoundTrip(q *http.Request) (*http.Response, error) {
	start := time.Now()
	result, e := t.base.RoundTrip(q)
	var code string

	if e == nil {
		code = strconv.Itoa(result.StatusCode)
	} else {
		code = failureCode(e)
	}

	t.observer.With(
		prometheus.Labels{
			constant.MetricHostLabel:   host(q.URL),
			constant.MetricCodeLabel:   code,
			constant.MetricMethodLabel: q.Method,
		},
	).Observe(time.Since(start).Seconds())

	return result, e
}
