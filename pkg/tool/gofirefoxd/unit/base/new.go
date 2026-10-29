package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Client = func() face.Client { return c }
	v := service.New(o)
	t.Cleanup(v.Close)
	r := memory.New()
	h := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gofirefoxd.Mount(v, r, mock_recorder.New(), g)
		},
	)
	result := &Server{
		MockClient: c,
		Service:    v,
		Reporter:   r,
		Directory:  o.DownloadDirectory,
		Server:     h,
	}
	t.Cleanup(result.Close)

	return result
}
