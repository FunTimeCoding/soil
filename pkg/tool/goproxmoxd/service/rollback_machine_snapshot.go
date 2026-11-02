package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) RollbackMachineSnapshot(
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

			task, e := c.RollbackMachineSnapshot(vm, name)

			if e != nil {
				return "", e
			}

			return string(task.UPID), nil
		},
	)
}
