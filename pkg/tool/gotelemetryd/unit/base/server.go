package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Server struct {
	Store    *store.Store
	Recorder *mock_recorder.Recorder
	*model_context_server.Server
}
