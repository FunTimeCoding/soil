package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	source := mock_client.New(t)
	result := &Tester{
		Source:   source,
		Reporter: memory.New(),
		client:   source,
		litePath: filepath.Join(t.TempDir(), "gonetboxd.db"),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
