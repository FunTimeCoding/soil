package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
)

type Daemon struct {
	runner  *runner.Runner
	service *service.Service
}
