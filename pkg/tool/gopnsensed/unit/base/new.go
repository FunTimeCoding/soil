package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"testing"
)

func New(
	t *testing.T,
	source face.OpnsenseClient,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{Reporter: memory.New(), source: source, t: t}

	for _, f := range o {
		f(result)
	}

	return result
}
