package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/daemon"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func connect(
	t *testing.T,
	d *daemon.Daemon,
	token string,
) *client.Client {
	t.Helper()
	s := model_context_server.Mount(t, d, memory.New())
	result := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		token,
	)
	t.Cleanup(result.Close)

	return result
}
