package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()

	return service.New(c)
}
