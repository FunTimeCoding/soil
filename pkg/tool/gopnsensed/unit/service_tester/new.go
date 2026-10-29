package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
	"testing"
)

func New(
	t *testing.T,
	c face.OpnsenseSource,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Source = func() face.OpnsenseSource { return c }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
