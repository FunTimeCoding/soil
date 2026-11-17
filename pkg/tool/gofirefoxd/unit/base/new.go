package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New(t)
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Client = func() face.FirefoxClient { return c }
	d := daemon.New(o)
	t.Cleanup(d.Close)
	r := memory.New()
	result := &Server{
		Source:    c,
		Daemon:    d,
		Reporter:  r,
		Directory: o.DownloadDirectory,
		t:         t,
		Server:    model_context_server.Mount(t, d, r),
	}
	t.Cleanup(result.Close)

	return result
}
