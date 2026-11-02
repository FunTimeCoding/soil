package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
)

func ContainerService(c *mock_client.Client) *service.Service {
	result := service.New(
		containerOption(),
		logger.New(context.Background()),
		memory.New(),
	)
	result.Pool().Set("test", c)

	return result
}
