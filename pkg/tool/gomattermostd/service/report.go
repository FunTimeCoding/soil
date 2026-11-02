package service

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (s *Service) report(failures []error) {
	for _, e := range failures {
		if classify.Reportable(e) {
			s.reporter.CaptureException(e)

			continue
		}

		s.logger.Structured("flush deferred", "error", e.Error())
	}
}
