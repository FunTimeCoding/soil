package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"sync"
)

type Worker struct {
	service  *service.Service
	logger   *logger.Logger
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
	running  sync.WaitGroup
}
