package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"testing"
)

type Server struct {
	*service_tester.Tester
	Reporter *memory.Memory
	served   *model_context_server.Server
	t        *testing.T
}
