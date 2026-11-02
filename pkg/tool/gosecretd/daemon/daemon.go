package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
)

type Daemon struct {
	source  face.SecretSource
	service *service.Service
	watch   *watch.Watch
}
