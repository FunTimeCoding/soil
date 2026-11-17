package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"testing"
)

func New(
	t *testing.T,
	path string,
) *Server {
	t.Helper()
	o := option.New()
	o.Database = path
	o.Password = "secret"
	r := memory.New()

	return &Server{
		Path:     path,
		Reporter: r,
		daemon:   daemon.New(o, logger.New(context.Background()), r),
		t:        t,
	}
}
