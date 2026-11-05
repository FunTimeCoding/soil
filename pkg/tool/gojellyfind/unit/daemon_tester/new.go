package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
	"testing"
)

func New(
	t *testing.T,
	c face.JellyfinClient,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.JellyfinClient { return c }
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
