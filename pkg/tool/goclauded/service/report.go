package service

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (s *Service) report(e error) {
	if e != nil && classify.Reportable(e) {
		s.reporter.CaptureException(e)
	}
}
