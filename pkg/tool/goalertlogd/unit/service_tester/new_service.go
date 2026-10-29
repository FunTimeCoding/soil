package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"testing"
)

func NewService(
	t *testing.T,
	path string,
	a face.Client,
) *service.Service {
	t.Helper()
	result := service.New(
		NewOption(path, a),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(result.Close)

	return result
}
