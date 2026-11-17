package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"io"
	"testing"
)

func New(
	t *testing.T,
	stream func(
		string,
		io.Writer,
		...string,
	) face.Stream,
	o ...Option,
) *Tester {
	t.Helper()
	p := &setup{reporter: memory.New()}

	for _, f := range o {
		f(p)
	}

	n := runner.New(t.TempDir(), stream, logger.New(t.Context()), p.reporter)
	t.Cleanup(n.Close)

	return &Tester{Runner: n}
}
