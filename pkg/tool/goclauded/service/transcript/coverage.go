package transcript

import "github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"

func (s *Service) Coverage() []*coverage.Server {
	s.reports.RLock()
	held := s.heldCoverage
	s.reports.RUnlock()

	if held != nil {
		return held
	}

	result := s.computeCoverage()
	s.reports.Lock()
	s.heldCoverage = result
	s.reports.Unlock()

	return result
}
