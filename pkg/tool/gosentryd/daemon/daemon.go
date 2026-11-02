package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
)

type Daemon struct {
	client  face.SentryClient
	guard   *peer.Guard
	service *service.Service
}
