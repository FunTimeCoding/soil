package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) ResetMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
) (string, error) {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.ResetMachine(vm)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
