package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
)

type Server struct {
	*service_tester.Tester
	served *model_context_server.Server
	t      *testing.T
}
