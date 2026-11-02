package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListStorageContent(
	instance string,
	node string,
	storage string,
) ([]*proxmox.StorageContent, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) ([]*proxmox.StorageContent, error) {
			n, e := c.Node(node)

			if e != nil {
				return nil, e
			}

			s2, e := c.Storage(n, storage)

			if e != nil {
				return nil, e
			}

			return c.StorageContent(s2)
		},
	)
}
