package service

import (
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory:  i,
		clients:    make(map[string]face.ProxmoxClient),
		sshClients: make(map[string]*ssh.Client),
	}
}
