package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

func IntentService(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
	machines ...*client.VirtualMachine,
) (*service.Service, *netbox.Client) {
	t.Helper()
	n := netbox.New()
	n.VirtualMachineList = append(n.VirtualMachineList, machines...)
	o := containerOption()
	o.Netbox = func() face.Client { return n }
	result := service.New(o, logger.New(t.Context()), r)
	result.Pool().Set("test", c)
	t.Cleanup(result.Close)

	return result, n
}
