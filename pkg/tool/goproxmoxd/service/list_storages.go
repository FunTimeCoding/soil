package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListStorages(
	c face.ProxmoxClient,
	node string,
) (proxmox.Storages, error) {
	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Storages(n)
}
