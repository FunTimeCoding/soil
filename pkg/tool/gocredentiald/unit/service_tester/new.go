package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
	"time"
)

func New(
	t *testing.T,
	path string,
	o ...Option,
) *service.Service {
	t.Helper()
	s := &Tester{now: time.Now, reporter: memory.New()}

	for _, f := range o {
		f(s)
	}

	return service.New(
		keepass.MustNew(path, "secret"),
		path,
		"secret",
		s.revealed,
		s.now,
		logger.New(context.Background()),
		s.reporter,
	)
}
