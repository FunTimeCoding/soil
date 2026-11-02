package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.SublimeSource { return c }
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
