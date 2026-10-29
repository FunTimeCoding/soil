package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			if s.watcher != nil {
				s.watcher.Close()
				s.sweeper.Close()
				s.connector.Close()
			}

			if s.monitor != nil {
				s.monitor.Close()
			}

			s.guard.Close()
			s.mattermost.Close()
			relational.CloseMapper(s.database)
		},
	)
}
