package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Service) QueuedCounts(
	sessions []session.Session,
) (map[string]int64, error) {
	return s.store.QueuedCounts(sessions)
}
