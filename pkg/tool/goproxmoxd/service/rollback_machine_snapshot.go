package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) RollbackMachineSnapshot(
	c face.ProxmoxClient,
	identifier int,
	node string,
	name string,
) (string, error) {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.RollbackMachineSnapshot(vm, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
