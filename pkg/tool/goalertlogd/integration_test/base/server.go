package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

type Server struct {
	Store         *store.Store
	Worker        *worker.Worker
	MockClient    *mock_client.Client
	ContextServer *model_context_server.Server
}
