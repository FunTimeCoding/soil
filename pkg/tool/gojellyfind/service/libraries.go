package service

import "github.com/funtimecoding/soil/pkg/jellyfin/library"

func (s *Service) Libraries() ([]*library.Library, error) {
	return s.source.Libraries()
}
