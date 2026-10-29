package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

type Server struct {
	t       *testing.T
	service *service.Service
	*model_context_server.Server
}
