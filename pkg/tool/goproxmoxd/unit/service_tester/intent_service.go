package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func IntentService(
	t *testing.T,
	c *mock_client.Client,
	n *netbox.Client,
	r *memory.Memory,
) *service.Service {
	t.Helper()
	result := ContainerService(c).WithNetbox(n, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
