package base

import (
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
)

type Server struct {
	Upstream      *mattermost_client_tester.Tester
	Service       *service.Service
	Reporter      *memory.Memory
	ContextServer *model_context_server.Server
}
