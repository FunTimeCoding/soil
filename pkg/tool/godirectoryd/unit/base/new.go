package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{
		Source:   service_tester.New(t).Source,
		Reporter: memory.New(),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
