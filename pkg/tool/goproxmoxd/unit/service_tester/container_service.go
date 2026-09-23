package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
)

func ContainerService(c *mock_client.Client) *service.Service {
	result := service.New(inventory.NewSingle("test"))
	result.SetClient("test", c)

	return result
}
