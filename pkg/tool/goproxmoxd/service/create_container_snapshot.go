package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) CreateContainerSnapshot(
	instance string,
	identifier int,
	node string,
	name string,
) (string, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (string, error) {
			ct, e := findContainer(c, identifier, node)

			if e != nil {
				return "", e
			}

			task, e := c.CreateContainerSnapshot(ct, name)

			if e != nil {
				return "", e
			}

			return string(task.UPID), nil
		},
	)
}
