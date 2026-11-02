package service

import "github.com/funtimecoding/soil/pkg/proxmox/face"

func (s *Service) DownloadLocator(
	instance string,
	node string,
	storage string,
	content string,
	filename string,
	l string,
) error {
	return s.attempt(
		instance,
		func(c face.ProxmoxClient) error {
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
		},
	)
}
