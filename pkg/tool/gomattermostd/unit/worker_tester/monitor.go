package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
)

func (o *Tester) Monitor(
	c *mock_client.Client,
	r *memory.Memory,
	channels ...string,
) *monitor.Monitor {
	o.t.Helper()
	p := option.New()
	p.Channels = channels

	return monitor.New(c, o.Guard(c, r), p, logger.New(o.t.Context()))
}
