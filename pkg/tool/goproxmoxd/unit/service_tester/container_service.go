package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func ContainerService(c *mock_client.Client) *service.Service {
	result := service.New(
		inventory.New[instance.Instance](
			instance.Instance{Name: "test", Host: "mock"},
		),
	)
	result.Pool().Set("test", c)

	return result
}
