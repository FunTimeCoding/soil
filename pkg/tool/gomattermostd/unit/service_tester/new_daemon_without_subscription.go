package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"path/filepath"
	"testing"
)

func NewDaemonWithoutSubscription(
	t *testing.T,
	c face.MattermostClient,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "mattermost.db")
	o.Mattermost = func() face.MattermostClient { return c }
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
