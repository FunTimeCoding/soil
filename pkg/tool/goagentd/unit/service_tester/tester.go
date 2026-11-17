package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
)

type Tester struct {
	Runner  *runner.Runner
	service *service.Service
}
