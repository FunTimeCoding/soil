package store_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"testing"
)

func NewService(t *testing.T) *service.Service {
	t.Helper()
	f := mock_client.New()
	t.Cleanup(f.Close)
	result := service.New(
		NewOption(t, f),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(result.Close)

	return result
}
