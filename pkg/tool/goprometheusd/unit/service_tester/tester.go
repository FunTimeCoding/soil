package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"testing"
	"time"
)

type Tester struct {
	instances map[string]*mock_client.Client
	Reporter  *memory.Memory
	Clock     func() time.Time
	service   *service.Service
	t         *testing.T
}
