package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func (s *Service) StopMachine(
	instance string,
	identifier int,
	node string,
) (string, error) {
	var name string
	upid, e := guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (string, error) {
			vm, f := findMachine(c, identifier, node)

			if f != nil {
				return "", f
			}

			task, f := c.StopMachine(vm)

			if f != nil {
				return "", f
			}

			name = vm.Name

			return string(task.UPID), nil
		},
	)

	if e != nil {
		return "", e
	}

	s.recordGuest(name, gonetboxdConstant.StatusOffline)

	return upid, nil
}
