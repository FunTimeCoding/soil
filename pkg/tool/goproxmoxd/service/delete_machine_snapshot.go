package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) DeleteMachineSnapshot(
	instance string,
	identifier int,
	node string,
	name string,
) (string, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (string, error) {
			vm, e := findMachine(c, identifier, node)

			if e != nil {
				return "", e
			}

			task, e := c.DeleteMachineSnapshot(vm, name)

			if e != nil {
				return "", e
			}

			return string(task.UPID), nil
		},
	)
}
