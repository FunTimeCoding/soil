package daemon

import (
	system "github.com/funtimecoding/soil/pkg/system/service"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/service"
)

func New() *Daemon {
	return &Daemon{service: service.New(system.New())}
}
