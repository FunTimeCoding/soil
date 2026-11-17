package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"testing"
)

type Tester struct {
	Source      *mock_client.Client
	Reporter    *memory.Memory
	client      face.NetBoxClient
	litePath    string
	powerIntent bool
	service     *service.Service
	daemon      *daemon.Daemon
	t           *testing.T
}
