package service_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
)

func Service(c *mock_client.Client) *service.Service {
	return service.New(c)
}
