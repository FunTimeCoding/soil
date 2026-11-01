package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
	"io"
	"testing"
)

func New(
	t *testing.T,
	stream func(string, io.Writer, ...string) face.Stream,
) *service.Service {
	t.Helper()

	return NewWith(t, stream, memory.New())
}
