package daemon

import "github.com/funtimecoding/soil/pkg/tool/goalpined/service"

type Daemon struct {
	service     *service.Service
	packageRoot string
}
