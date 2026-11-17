package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"testing"
)

type Tester struct {
	instances map[string]*mock_client.Client
	Reporter  *memory.Memory
	service   *service.Service
	t         *testing.T
}
