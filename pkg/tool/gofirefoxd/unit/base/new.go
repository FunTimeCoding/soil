package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Client = func() face.FirefoxClient { return c }
	d := daemon.New(o)
	t.Cleanup(d.Close)
	r := memory.New()
	h := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	result := &Server{
		MockClient: c,
		Daemon:     d,
		Reporter:   r,
		Directory:  o.DownloadDirectory,
		Server:     h,
	}
	t.Cleanup(result.Close)

	return result
}
