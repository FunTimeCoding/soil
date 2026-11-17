package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		instances: map[string]*mock_client.Client{"alfa": mock_client.New(t)},
		Reporter:  memory.New(),
		Clock:     time.Now,
		t:         t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
