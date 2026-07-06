package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) DeleteContainerSnapshot(
	c face.ProxmoxClient,
	identifier int,
	node string,
	name string,
) (string, error) {
	ct, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.DeleteContainerSnapshot(ct, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
