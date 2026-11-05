package daemon

import (
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
)

type Daemon struct {
	client  face.OpenbaoClient
	guard   *peer.Guard
	service *service.Service
	watch   *watch.Watch
}
