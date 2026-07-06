package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
	"sync"
)

type Service struct {
	inventory  *inventory.Inventory
	clients    map[string]*proxmox.Client
	sshClients map[string]*ssh.Client
	sessions   sync.Map
	mutex      sync.Mutex
}
