package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"testing"
)

type Tester struct {
	Netbox    *netbox.Client
	Reporter  *memory.Memory
	instances map[string]*Clients
	redial    bool
	dialled   []*sshMock.Client
	service   *service.Service
	t         *testing.T
}
