package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"testing"
	"time"
)

type Tester struct {
	Forge    *mock_client.Client
	Reporter *memory.Memory
	Clock    func() time.Time
	service  *service.Service
	t        *testing.T
}
