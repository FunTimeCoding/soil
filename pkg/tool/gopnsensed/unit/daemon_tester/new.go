package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"testing"
)

func New(
	t *testing.T,
	c face.OpnsenseSource,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.OpnsenseSource { return c }
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
