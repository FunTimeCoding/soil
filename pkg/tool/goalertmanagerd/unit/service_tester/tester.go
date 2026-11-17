package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"testing"
)

type Tester struct {
	Reporter  *memory.Memory
	instances map[string]*Clients
	service   *service.Service
	t         *testing.T
}
