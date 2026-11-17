package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "maintenance.db")
	r := memory.New()
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)
	result := &Server{
		Reporter: r,
		LitePath: o.LitePath,
		t:        t,
		Server:   model_context_server.Mount(t, d, r),
	}
	t.Cleanup(result.Close)

	return result
}
