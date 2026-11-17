package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/metric"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{Reporter: memory.New(), Metric: metric.New(), t: t}
	WithInstances("alfa")(result)

	for _, f := range o {
		f(result)
	}

	return result
}
