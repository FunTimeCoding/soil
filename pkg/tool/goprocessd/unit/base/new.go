package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func New(
	t *testing.T,
	procfile string,
) *Server {
	t.Helper()

	return &Server{Reporter: memory.New(), procfile: procfile, t: t}
}
