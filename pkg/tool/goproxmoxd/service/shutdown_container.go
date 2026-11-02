package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func (s *Service) ShutdownContainer(
	instance string,
	identifier int,
	node string,
) (string, error) {
	var name string
	upid, e := guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (string, error) {
			container, f := findContainer(c, identifier, node)

			if f != nil {
				return "", f
			}

			task, f := c.ShutdownContainer(container)

			if f != nil {
				return "", f
			}

			name = container.Name

			return string(task.UPID), nil
		},
	)

	if e != nil {
		return "", e
	}

	s.recordGuest(name, gonetboxdConstant.StatusOffline)

	return upid, nil
}
