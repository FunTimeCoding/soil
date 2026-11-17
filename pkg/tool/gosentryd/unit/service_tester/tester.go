package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
	"testing"
)

type Tester struct {
	Source   *mock_client.Client
	Reporter *memory.Memory
	service  *service.Service
	t        *testing.T
}
