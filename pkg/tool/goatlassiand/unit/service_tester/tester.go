package service_tester

import (
	confluenceMock "github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	jiraMock "github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"testing"
)

type Tester struct {
	Jira       *jiraMock.Client
	Confluence *confluenceMock.Client
	Reporter   *memory.Memory
	project    []string
	notifier   face.EventNotifier
	service    *service.Service
	daemon     *daemon.Daemon
	t          *testing.T
}
