package service

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"

func (s *Service) WithMonitor(m *monitor.Monitor) *Service {
	s.monitor = m

	return s
}
