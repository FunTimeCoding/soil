package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{Tester: service_tester.NewFixture(t, nil), t: t}

	for _, f := range o {
		f(result)
	}

	return result
}
