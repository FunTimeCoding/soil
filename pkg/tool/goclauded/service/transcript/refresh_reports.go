package transcript

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) RefreshReports() []error {
	servers := s.computeCoverage()
	report, e := s.computePermissions(constant.PermissionWindowDays)
	s.reports.Lock()
	defer s.reports.Unlock()
	s.heldCoverage = servers

	if e != nil {
		return []error{e}
	}

	s.heldPermissions = report

	return nil
}
