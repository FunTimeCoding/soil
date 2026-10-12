package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
)

func Service(c *mock_client.Client) *service.Service {
	return service.New(c, "test")
}
