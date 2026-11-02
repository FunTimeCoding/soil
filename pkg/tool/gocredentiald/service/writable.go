package service

import "github.com/funtimecoding/soil/pkg/errors/conflict"

func (s *Service) writable() error {
	s.refresh()

	if s.client.Changed() {
		return conflict.Format(
			"database changed on disk and cannot be reloaded - refusing to write",
		)
	}

	return nil
}
