package service_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
)

func Single(c *mock_client.Client) *service.Service {
	return Service(map[string]*mock_client.Client{constant.FixtureInstance: c})
}
