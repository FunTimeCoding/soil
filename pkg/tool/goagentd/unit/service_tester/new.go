package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"io"
	"testing"
)

func New(
	t *testing.T,
	stream func(string, io.Writer, ...string) face.Stream,
) *Tester {
	t.Helper()

	return NewWith(t, stream, memory.New())
}
