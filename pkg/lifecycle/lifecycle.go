package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle/worker"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

type Lifecycle struct {
	component []face.Component
	drainer   []*worker.Worker
	logger    *logger.Logger
	grace     time.Duration
}
