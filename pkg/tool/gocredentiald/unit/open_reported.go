package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
	"time"
)

func openReported(
	t *testing.T,
	path string,
	revealedField []string,
	r face.Reporter,
) *service.Service {
	t.Helper()

	return service.New(
		keepass.MustNew(path, "secret"),
		path,
		"secret",
		revealedField,
		time.Now,
		logger.New(context.Background()),
		r,
	)
}
