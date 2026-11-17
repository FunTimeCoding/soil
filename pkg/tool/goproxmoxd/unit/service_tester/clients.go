package service_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
)

type Clients struct {
	Hypervisor *mock_client.Client
	Shell      *sshMock.Client
}
