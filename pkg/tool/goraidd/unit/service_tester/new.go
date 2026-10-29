package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"testing"
)

func New(
	t *testing.T,
	logCachePath string,
	outputPath string,
) *service.Service {
	t.Helper()

	return NewWith(t, logCachePath, outputPath, NewParser(t), memory.New())
}
