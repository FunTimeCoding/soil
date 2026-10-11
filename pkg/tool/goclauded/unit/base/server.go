package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
)

type Server struct {
	*service_tester.Tester
	*model_context_server.Server
}
