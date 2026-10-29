package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
	"testing"
)

func New(
	t *testing.T,
	c face.JellyfinSource,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Source = func() face.JellyfinSource { return c }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
