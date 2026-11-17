package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Prometheus: mock_client.New(t),
		Reporter:   memory.New(),
		board:      FixtureBoard(),
		litePath:   filepath.Join(t.TempDir(), "godashboardd.db"),
		t:          t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
