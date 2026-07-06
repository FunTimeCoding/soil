package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

type Lifecycle struct {
	component []face.Component
	logger    *logger.Logger
	verbose   bool
}
