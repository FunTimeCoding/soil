package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Source:   mock_client.New(t),
		Reporter: memory.New(),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
