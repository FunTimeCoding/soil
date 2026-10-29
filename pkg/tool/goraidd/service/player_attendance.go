package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"
	"time"
)

func (s *Service) PlayerAttendance(since time.Time) []record.Attendance {
	return s.store.PlayerAttendance(since)
}
