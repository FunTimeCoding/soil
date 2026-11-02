package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
	"time"
)

func openAt(
	t *testing.T,
	path string,
	at time.Time,
) *service.Service {
	t.Helper()

	return service.New(
		keepass.MustNew(path, "secret"),
		path,
		"secret",
		nil,
		func() time.Time { return at },
		logger.New(context.Background()),
		memory.New(),
	)
}
