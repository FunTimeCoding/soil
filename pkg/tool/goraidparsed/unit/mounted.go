package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func mounted(
	t *testing.T,
	p face.TopStatsParser,
) *model_context_server.Server {
	t.Helper()
	o := option.New()
	o.TemplatePath = t.TempDir()
	o.OutputPath = t.TempDir()
	o.TopStats = func(string) face.TopStatsParser { return p }
	d := daemon.New(o, logger.New(t.Context()))
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			m *guard.Mux,
		) {
			d.Mount(memory.New(), mock_recorder.New(), m)
		},
	)
	t.Cleanup(result.Close)

	return result
}
