package outpost

import "github.com/funtimecoding/soil/pkg/log/logger"

type Collector struct {
	targets []*Target
	logger  *logger.Logger
}
