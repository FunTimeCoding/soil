package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
	"testing"
)

type Tester struct {
	Runner  *runner.Runner
	service *service.Service
	t       *testing.T
}
