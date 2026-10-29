package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/face"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Source = func() face.SublimeSource { return c }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
