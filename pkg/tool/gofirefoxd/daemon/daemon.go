package daemon

import (
	"github.com/funtimecoding/soil/pkg/firefox/face"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/service"
	"sync"
)

type Daemon struct {
	client    face.FirefoxClient
	service   *service.Service
	closeOnce sync.Once
}
