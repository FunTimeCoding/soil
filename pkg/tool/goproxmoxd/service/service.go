package service

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Service struct {
	inventory   *instance.Inventory
	pool        *pool.Pool[instance.Instance, *instance.Instance, proxmox.ProxmoxClient]
	shells      *pool.Pool[instance.Instance, *instance.Instance, ssh.SSHClient]
	netbox      netbox.Client
	netboxGuard *peer.Guard
	logger      *logger.Logger
	reporter    soilFace.Reporter
}
