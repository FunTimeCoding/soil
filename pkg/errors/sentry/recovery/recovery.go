package recovery

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

type Recovery struct {
	logger   *logger.Logger
	reporter face.Reporter
}
