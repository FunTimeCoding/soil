package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) client(instance string) (face.ProxmoxClient, error) {
	return s.pool.Client(instance)
}
