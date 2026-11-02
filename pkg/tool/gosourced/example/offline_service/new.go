package offline_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func New() *service.Service {
	o := option.New()
	o.Memory = func() face.Client { return mock_client.New() }

	return service.New(o, logger.New(context.Background()), memory.New())
}
