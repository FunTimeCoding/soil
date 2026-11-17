package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"testing"
)

func New(
	t *testing.T,
	p face.TopStatsParser,
) *Server {
	t.Helper()
	o := option.New()
	o.TemplatePath = t.TempDir()
	o.OutputPath = t.TempDir()
	o.TopStats = func(string) face.TopStatsParser { return p }

	return &Server{Reporter: memory.New(), option: o, t: t}
}
