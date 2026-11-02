package base

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	jiraMock "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/daemon"
)

type Server struct {
	MockClient *mock_client.Client
	Jira       *jiraMock.Client
	Daemon     *daemon.Daemon
	*model_context_server.Server
}
