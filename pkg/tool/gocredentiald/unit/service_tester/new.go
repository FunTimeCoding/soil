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
) *Tester {
	t.Helper()
	result := &Tester{Reporter: memory.New(), now: time.Now}

	for _, f := range o {
		f(result)
	}

	result.service = service.New(
		keepass.MustNew(path, "secret"),
		path,
		"secret",
		result.revealed,
		result.now,
		logger.New(context.Background()),
		result.Reporter,
	)

	return result
}
