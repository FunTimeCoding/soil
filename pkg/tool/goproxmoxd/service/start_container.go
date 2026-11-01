package service

import netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"

func (s *Service) StartContainer(
	instance string,
	identifier int,
	node string,
) (string, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return "", clientFail
	}

	container, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.StartContainer(container)

	if e != nil {
		return "", e
	}

	s.recordGuest(container.Name, netbox.StatusActive)

	return string(task.UPID), nil
}
