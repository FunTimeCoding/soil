package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) StopContainer(
	c face.ProxmoxClient,
	identifier int,
	node string,
) (string, error) {
	container, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.StopContainer(container)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
