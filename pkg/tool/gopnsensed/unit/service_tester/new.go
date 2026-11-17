package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"testing"
)

func New(
	t *testing.T,
	c face.OpnsenseClient,
) *Tester {
	t.Helper()

	return &Tester{Reporter: memory.New(), source: c, t: t}
}
