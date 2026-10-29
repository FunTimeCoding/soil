package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
)

type Server struct {
	Store face.Claims
	*model_context_server.Server
}
