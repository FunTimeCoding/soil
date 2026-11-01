package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
	"io"
	"testing"
)

func NewWith(
	t *testing.T,
	stream func(string, io.Writer, ...string) face.Stream,
	r face.Reporter,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Workspace = t.TempDir()
	o.Stream = stream
	result := service.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
