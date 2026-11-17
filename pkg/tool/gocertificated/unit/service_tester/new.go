package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Forge:    mock_client.New(t),
		Reporter: memory.New(),
		Clock:    time.Now,
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
