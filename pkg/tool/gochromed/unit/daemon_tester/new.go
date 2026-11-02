package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Source = func() face.ChromiumSource { return c }
	result := daemon.New(o)
	t.Cleanup(result.Close)

	return result
}
