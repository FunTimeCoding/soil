package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
	"time"
)

func openRevealed(
	t *testing.T,
	path string,
	revealedField []string,
) *service.Service {
	t.Helper()

	return service.New(
		path,
		"secret",
		revealedField,
		time.Now,
		logger.New(context.Background()),
	)
}
