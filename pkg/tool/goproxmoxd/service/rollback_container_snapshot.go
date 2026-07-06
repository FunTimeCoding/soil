package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) RollbackContainerSnapshot(
	c face.ProxmoxClient,
	identifier int,
	node string,
	name string,
) (string, error) {
	ct, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.RollbackContainerSnapshot(ct, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
