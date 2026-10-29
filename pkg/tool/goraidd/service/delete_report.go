package service

import "github.com/funtimecoding/soil/pkg/system"

func (s *Service) DeleteReport(name string) error {
	path, e := s.reportFile(name)

	if e != nil {
		return e
	}

	system.RemoveFile(path)

	return nil
}
