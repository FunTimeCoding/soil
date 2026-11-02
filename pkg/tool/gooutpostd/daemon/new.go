package daemon

import (
	"github.com/funtimecoding/soil/pkg/system/service"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/service"
)

func New() *Daemon {
	return &Daemon{service: outpost.New(service.New())}
}
