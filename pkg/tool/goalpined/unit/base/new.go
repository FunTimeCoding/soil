package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()

	return &Server{
		Tester:   service_tester.New(t),
		Reporter: memory.New(),
		t:        t,
	}
}
