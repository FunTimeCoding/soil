package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"testing"
)

type Tester struct {
	Mattermost *mock_client.Client
	Reporter   *memory.Memory
	Path       string
	client     chat.MattermostClient
	connector  goclauded.Client
	monitor    bool
	service    *service.Service
	daemon     *daemon.Daemon
	t          *testing.T
}
