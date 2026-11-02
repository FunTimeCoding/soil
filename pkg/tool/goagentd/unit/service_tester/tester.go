package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
)

type Tester struct {
	Service *service.Service
	Runner  *runner.Runner
}
