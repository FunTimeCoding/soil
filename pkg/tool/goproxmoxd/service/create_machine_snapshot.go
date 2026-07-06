package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) CreateMachineSnapshot(
	c face.ProxmoxClient,
	identifier int,
	node string,
	name string,
) (string, error) {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.CreateMachineSnapshot(vm, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
