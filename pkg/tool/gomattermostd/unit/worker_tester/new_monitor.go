package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"testing"
)

func NewMonitor(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
	channels ...string,
) *monitor.Monitor {
	t.Helper()
	o := option.New()
	o.Channels = channels

	return monitor.New(c, NewGuard(t, c, r), o, logger.New(t.Context()))
}
