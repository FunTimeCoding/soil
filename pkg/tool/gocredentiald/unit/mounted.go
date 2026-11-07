package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func mounted(
	t *testing.T,
	path string,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	o := option.New()
	o.Database = path
	o.Password = "secret"
	d := daemon.New(o, logger.New(context.Background()), r)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
