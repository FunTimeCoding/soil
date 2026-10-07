package outpost

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/target"
)

type Collector struct {
	targets []*target.Target
	logger  *logger.Logger
}
