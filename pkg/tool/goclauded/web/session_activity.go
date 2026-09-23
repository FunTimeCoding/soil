package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func sessionActivity(s *session.Session) time.Time {
	if s.LastActiveAt.After(s.LastSeen) {
		return s.LastActiveAt
	}

	return s.LastSeen
}
