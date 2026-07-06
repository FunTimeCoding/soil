package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory:  i,
		clients:    make(map[string]*proxmox.Client),
		sshClients: make(map[string]*ssh.Client),
	}
}
