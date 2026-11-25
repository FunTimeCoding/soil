package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	gitlabMock "github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"testing"
)

type Tester struct {
	instances map[string]*gitlabMock.Client
	connector *mock_client.Client
	Reporter  *memory.Memory
	Path      string
	service   *service.Service
	t         *testing.T
}
