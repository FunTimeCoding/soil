package service_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/ladybird/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
	"testing"
)

type Tester struct {
	Client   *mock_client.Client
	Reporter *memory.Memory
	service  *service.Service
	t        *testing.T
}
