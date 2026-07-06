package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

type Builder struct {
	name         string
	version      string
	instructions string
	logger       *logger.Logger
	verbose      bool
	resources    bool
	recorder     face.Recorder
}
