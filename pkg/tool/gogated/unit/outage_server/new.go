package outage_server

import (
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := mock_client.New("")
	c.Failure = outage_tester.RefusedDial()
	s := base.New(t, base.WithDirectory(c), base.WithReporter(r))

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		constant.DirectoryPeer,
	).Unavailable
}
