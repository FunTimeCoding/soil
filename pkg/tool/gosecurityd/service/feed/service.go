package feed

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
)

type Service struct {
	ansible face.Client
	guard   *peer.Guard
	system  *system.Service
	logger  *logger.Logger
}
