package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
	"testing"
)

func New(
	t *testing.T,
	c face.JellyfinSource,
) *service.Service {
	t.Helper()

	return service.New(c)
}
