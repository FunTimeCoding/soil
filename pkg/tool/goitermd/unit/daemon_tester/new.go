package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.ItermSource { return c }
	o.KeyInterval = 0
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
