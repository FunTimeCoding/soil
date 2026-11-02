package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
	"testing"
)

func New(
	t *testing.T,
	c face.OpnsenseSource,
) *service.Service {
	t.Helper()

	return service.New(c)
}
