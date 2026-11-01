package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func activity(s *session.Session) time.Time {
	if s.LastActiveAt.After(s.LastSeen) {
		return s.LastActiveAt
	}

	return s.LastSeen
}
