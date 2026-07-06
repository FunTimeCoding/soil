package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListStorageContent(
	c face.ProxmoxClient,
	node string,
	storage string,
) ([]*proxmox.StorageContent, error) {
	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	s2, e := c.Storage(n, storage)

	if e != nil {
		return nil, e
	}

	return c.StorageContent(s2)
}
