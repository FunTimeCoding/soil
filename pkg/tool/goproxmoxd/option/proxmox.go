package option

import (
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Proxmox struct {
	Address       string
	ServiceTokens []string
	MetricAddress string
	Inventory     *instance.Inventory
	Hypervisor    func(*instance.Instance) (proxmox.ProxmoxClient, error)
	Shell         func(*instance.Instance) (ssh.SSHClient, error)
	Netbox        func() face.Client
}
