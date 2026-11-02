package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"
)

type Daemon struct {
	supervisor *supervisor.Supervisor
	service    *service.Service
}
