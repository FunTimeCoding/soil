package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Source = func() face.ItermSource { return c }
	o.KeyInterval = 0
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
