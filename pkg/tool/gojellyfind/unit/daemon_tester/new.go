package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
	"testing"
)

func New(
	t *testing.T,
	c face.JellyfinSource,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.JellyfinSource { return c }
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
