package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
	"io"
	"testing"
)

func NewWith(
	t *testing.T,
	stream func(string, io.Writer, ...string) face.Stream,
	r face.Reporter,
) *Tester {
	t.Helper()
	n := runner.New(t.TempDir(), stream, logger.New(t.Context()), r)
	t.Cleanup(n.Close)

	return &Tester{Service: service.New(n), Runner: n}
}
