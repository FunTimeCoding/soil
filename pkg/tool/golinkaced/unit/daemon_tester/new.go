package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/linkace/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/daemon"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.LinkAceClient { return c }
	result := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
