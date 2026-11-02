package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			if s.watcher != nil {
				s.watcher.Close()

				for root, events := range s.watcher.Drain() {
					s.report(s.flush(root, events))
				}

				s.rosterGuard.Close()
				s.connector.Close()
			}

			s.guard.Close()
			s.mattermost.Close()
			relational.CloseMapper(s.database)
		},
	)
}
