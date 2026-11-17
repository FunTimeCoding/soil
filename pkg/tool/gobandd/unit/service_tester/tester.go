package service_tester

import (
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

type Tester struct {
	Netbox    *netbox.Client
	Reporter  *memory.Memory
	instances map[string]*mock_client.Client
	service   *service.Service
	t         *testing.T
}
