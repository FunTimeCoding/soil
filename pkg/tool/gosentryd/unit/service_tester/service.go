package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
	"testing"
)

func Service(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()

	return Shared(t, c, memory.New())
}
