package service

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"

func (s *Service) PlayerAttendance() []record.Attendance {
	return s.store.PlayerAttendance(s.clock().AddDate(0, -1, 0))
}
