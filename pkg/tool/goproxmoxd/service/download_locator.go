package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (s *Service) DownloadLocator(
	c face.ProxmoxClient,
	node string,
	storage string,
	content string,
	filename string,
	l string,
) error {
	n, e := c.Node(node)

	if e != nil {
		return e
	}

	st, e := c.Storage(n, storage)

	if e != nil {
		return e
	}

	task, e := c.DownloadLocator(st, content, filename, l)

	if e != nil {
		return e
	}

	return c.WaitForTask(task, 600)
}
