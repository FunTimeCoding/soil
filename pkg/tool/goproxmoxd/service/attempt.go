package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) attempt(
	instance string,
	f func(face.ProxmoxClient) error,
) error {
	_, e := guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (bool, error) {
			return true, f(c)
		},
	)

	return e
}
