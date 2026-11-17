package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{Tester: service_tester.New(t), t: t}

	for _, f := range o {
		f(result)
	}

	return result
}
