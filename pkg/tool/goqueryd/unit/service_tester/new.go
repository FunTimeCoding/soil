package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

func New(
	t *testing.T,
	o *option.Query,
	r face.Reporter,
) *service.Service {
	t.Helper()
	result := service.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
